package handlers

import (
	"auth-repo/authentication"
	"auth-repo/config"
)

type Handler struct {
	conf    *config.Config
	service authentication.Service
}

func NewHandler(conf *config.Config, service authentication.Service) *Handler {
	return &Handler{
		conf:    conf,
		service: service,
	}
}
