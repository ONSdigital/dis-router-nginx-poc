package middleware

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type Redirect struct {
	Path     string `json:"path"`
	Redirect string `json:"redirect"`
	Type     string `json:"type"`
}

const (
	RedirectPermenantKey = "permenant"
	RedirectTemporaryKey = "temporary"
)

func RedirectHandler(redisClient *redis.Client) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			slog.Info("checking if redirect for uri", "uri", req.RequestURI)

			val, err := redisClient.JSONGet(req.Context(), req.RequestURI, "$").Result()
			if err != nil {
				slog.Error("error getting from redis", "err", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if val != "" {
				var result []Redirect
				err = json.Unmarshal([]byte(val), &result)
				if err != nil {
					slog.Error("error unmarshalling json", "err", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				var status int

				switch result[0].Type {
				case RedirectPermenantKey:
					status = http.StatusPermanentRedirect
				case RedirectTemporaryKey:
					status = http.StatusTemporaryRedirect
				default:
					slog.Error("unknown redirect type", "type", result[0].Type)
					status = http.StatusTemporaryRedirect
				}

				http.Redirect(w, req, result[0].Redirect, status)
			} else {
				slog.Info("no redirect for uri, proxying", "uri", req.RequestURI)
				// No redirect, allow proxy
				h.ServeHTTP(w, req)
			}
			return
		})
	}
}
