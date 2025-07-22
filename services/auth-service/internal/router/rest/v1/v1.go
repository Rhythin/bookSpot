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
		r.Post("/logout", eh(handler.Logout))
		r.Get("/", eh(handler.GetUsers))
		r.Get("/{userID}", eh(handler.GetUser))
		r.Put("/{userID}", eh(handler.UpdateUser))
		r.Delete("/{userID}", eh(handler.DeleteUser))
		r.Post("/reset-password", eh(handler.ResetPassword))
	})

	r.Route("/token", func(r chi.Router) {
		r.Get("/", eh(handler.GetToken)) // query params: tempToken
		r.Post("/revoke", eh(handler.RevokeToken))
		r.Post("/refresh", eh(handler.RefreshToken))
	})

	return r
}
