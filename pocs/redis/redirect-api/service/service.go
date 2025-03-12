package service

import (
	"github.com/ONSdigital/dis-router-nginx-poc/redis/redirect-api/api"
	"github.com/ONSdigital/dis-router-nginx-poc/redis/redirect-api/redis"
)

type Service struct {
	server Server
}

func (s *Service) Run() error {
	redisClient := redis.NewRedisClient()

	s.server = Server{
		Name:    "proxyServer",
		Addr:    ":3003",
		Handler: api.NewAPIHandler(redisClient),
	}

	s.server.Start()

	return nil
}

func (s *Service) Shutdown() {
	s.server.Stop()
}
