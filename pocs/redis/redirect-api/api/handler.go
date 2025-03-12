package api

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func NewAPIHandler(redisClient *redis.Client) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /redirects", addRedirect(redisClient))
	return mux
}

type Redirect struct {
	Path     string `json:"path"`
	Redirect string `json:"redirect"`
	Type     string `json:"type"`
}

const (
	RedirectPermenantKey = "permenant"
	RedirectTemporaryKey = "temporary"
)

func addRedirect(redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		slog.Info("received add redirect request")
		body, err := io.ReadAll(req.Body)
		if err != nil {
			slog.Error("error reading body")
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		var payload []Redirect
		if err := json.Unmarshal(body, &payload); err != nil {
			slog.Error("error unmarshalling body")
			http.Error(w, "invalid json in request", http.StatusBadRequest)
			return
		}

		// TODO VALIDATE INPUTS!!!!

		for _, item := range payload {
			var status int
			switch item.Type {
			case RedirectPermenantKey:
				status = http.StatusPermanentRedirect
			case RedirectTemporaryKey:
				status = http.StatusTemporaryRedirect
			default:
				status = http.StatusTemporaryRedirect
			}

			itemBytes, err := json.Marshal(item)
			if err != nil {
				slog.Error("error marshalling item to json")
				http.Error(w, "invalid json in request", http.StatusInternalServerError)
				return
			}

			slog.Debug("adding redirect", "path", item.Path, "redirect", item.Redirect, "type", item.Type, "status", status)
			err = redisClient.JSONSet(req.Context(), item.Path, "$", itemBytes).Err()
			if err != nil {
				slog.Error("error setting redirect in redis", "err", err)
				http.Error(w, "error saving redirect", http.StatusInternalServerError)
				return
			}
		}

		w.WriteHeader(http.StatusCreated)
	}
}
