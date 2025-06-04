package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/rhythin/bookspot/auth-service/internal/handler/rest"
	"github.com/rhythin/bookspot/services/shared/errhandler"
)

func NewRouter(handler rest.Handler) chi.Router {
	eh := errhandler.HttpErrorHandler
	r := chi.NewRouter()

	r.Route("/user", func(r chi.Router) {
		r.Post("/register", eh(handler.Register))
		r.Post("/login", eh(handler.Login))
	})

	return r
}
