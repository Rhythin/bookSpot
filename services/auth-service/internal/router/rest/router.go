package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/riandyrn/otelchi"
	_ "github.com/rhythin/bookspot/auth-service/docs"
	"github.com/rhythin/bookspot/auth-service/internal/handler/rest"
	v1 "github.com/rhythin/bookspot/auth-service/internal/router/rest/v1"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NewRouter(handler rest.Handler) chi.Router {

	r := chi.NewRouter()

	// middleware
	r.Use(otelchi.Middleware("auth-service"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// routes
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
	))
	r.Route("/v1", func(r chi.Router) {
		v1.NewRouter(handler)
	})

	return r
}
