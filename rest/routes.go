package rest

import (
	"auth-repo/rest/middlewares"
	"net/http"
)

func (s *Server) initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"POST /register",
		manager.With(
			http.HandlerFunc(s.handler.Registration),
		),
	)

	mux.Handle(
		"GET /hello",
		manager.With(
			http.HandlerFunc(s.handler.Hello),
		),
	)
}
