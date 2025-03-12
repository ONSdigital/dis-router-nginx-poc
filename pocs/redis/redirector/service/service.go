package service

import (
	"github.com/ONSdigital/dis-router-nginx-poc/redis/redirector/router"
)

type Service struct {
	server Server
}

func (s *Service) Run() error {
	s.server = Server{
		Name:    "proxyServer",
		Addr:    ":3000",
		Handler: router.New(),
	}

	s.server.Start()

	return nil
}

func (s *Service) Shutdown() {
	s.server.Stop()
}
