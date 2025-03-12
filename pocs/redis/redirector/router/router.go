package router

import (
	"net/http"

	"github.com/ONSdigital/dis-router-nginx-poc/redis/redirector/middleware"
	"github.com/ONSdigital/dis-router-nginx-poc/redis/redirector/proxy"
	"github.com/ONSdigital/dis-router-nginx-poc/redis/redirector/redis"

	"github.com/justinas/alice"
)

func New() http.Handler {

	router := http.NewServeMux()

	redisClient := redis.NewRedisClient()

	//slog.Info("Adding default redirect to redis", "key", "/bacon")
	//redisClient.JSONSet(context.TODO(), "/bacon", "$", `{ "redirect": "/sausages", "type": "permenant"}`)

	middleware := []alice.Constructor{
		middleware.RedirectHandler(redisClient),
	}

	proxyHandler, err := proxy.NewProxyHandler("http://fakebabbage:3002")
	if err != nil {
		panic(err) // TODO handle error properly
	}

	router.Handle("/", proxyHandler)

	newAlice := alice.New(middleware...).Then(router)

	return newAlice
}
