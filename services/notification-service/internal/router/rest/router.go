package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rhythin/bookspot/notification-service/internal/handler"
	v1 "github.com/rhythin/bookspot/notification-service/internal/router/rest/v1"
)

func GetRouter(handler handler.Handler) chi.Router {

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// API routes
	r.Route("/v1", func(r chi.Router) {
		v1.NewRouter(handler.V1)
	})

	return r
}
