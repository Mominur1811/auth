package rest

import (
	"auth-repo/config"
	"auth-repo/rest/handlers"
	"auth-repo/rest/middlewares"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
)

type Server struct {
	handler *handlers.Handler
	conf    *config.Config
	Wg      sync.WaitGroup
}

func NewServer(conf *config.Config, handler *handlers.Handler) *Server {
	return &Server{
		handler: handler,
		conf:    conf,
	}
}

func (s *Server) Start() {

	manager := middlewares.NewManager()

	manager.Use(
		middlewares.Recover,
		middlewares.Logger,
	)
	mux := http.NewServeMux()
	s.initRoutes(mux, manager)

	handler := middlewares.EnableCors(mux)

	s.Wg.Add(1)

	go func() {
		defer s.Wg.Done()
		addr := fmt.Sprintf(":%d", s.conf.HttpPort)
		slog.Info(fmt.Sprintf("Starting server on %s", addr))

		if err := http.ListenAndServe(addr, handler); err != nil {
			slog.Error(err.Error())
		}
	}()
}
