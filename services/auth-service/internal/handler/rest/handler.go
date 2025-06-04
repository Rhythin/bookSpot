package rest

import (
	"net/http"

	"github.com/rhythin/bookspot/auth-service/internal/service"
)

type Handler interface {
	Register(w http.ResponseWriter, r *http.Request) (err error)
	Login(w http.ResponseWriter, r *http.Request) (err error)
}

type handler struct {
	service service.Service
}

func New(service service.Service) Handler {
	return &handler{
		service: service,
	}
}
