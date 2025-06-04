package v1

import (
	"github.com/go-chi/chi/v5"
	v1 "github.com/rhythin/bookspot/books-service/internal/handler/v1"
	errhandler "github.com/rhythin/bookspot/services/shared/errhandler"
)

func NewRouter(handler v1.HandlerV1) chi.Router {
	eh := errhandler.HttpErrorHandler
	r := chi.NewRouter()

	r.Route("/book", func(r chi.Router) {
		r.Get("/", eh(handler.GetBooks))
	})

	return r
}
