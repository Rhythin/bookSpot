package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rhythin/bookspot/auth-service/internal/handler/rest"
	v1 "github.com/rhythin/bookspot/auth-service/internal/router/rest/v1"
)

func NewRouter(handler rest.Handler) chi.Router {

	r := chi.NewRouter()

	// middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// routes
	r.Route("/v1", func(r chi.Router) {
		v1.NewRouter(handler)
	})

	return r
}
