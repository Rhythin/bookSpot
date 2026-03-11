package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/rhythin/bookspot/notification-service/docs"
	"github.com/rhythin/bookspot/notification-service/internal/handler"
	v1 "github.com/rhythin/bookspot/notification-service/internal/router/rest/v1"
	"github.com/rhythin/bookspot/services/shared/jwt_auth"
	"github.com/riandyrn/otelchi"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func GetRouter(handler handler.Handler, tokenizer jwt_auth.Tokenizer) chi.Router {

	r := chi.NewRouter()

	// Swagger route
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// jwt middleware
	authMw := jwt_auth.NewMiddleware(tokenizer)

	// Middleware
	r.Use(otelchi.Middleware("notification-service"))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(authMw.Authenticate)

	// API routes
	r.Route("/v1", func(r chi.Router) {
		v1.NewRouter(handler.V1)
	})

	return r
}
