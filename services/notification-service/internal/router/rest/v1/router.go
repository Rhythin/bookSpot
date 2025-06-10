package v1

import (
	"github.com/go-chi/chi/v5"
	v1 "github.com/rhythin/bookspot/notification-service/internal/handler/v1"
	errhandler "github.com/rhythin/bookspot/services/shared/errhandler"
)

func NewRouter(handler v1.HandlerV1) chi.Router {
	eh := errhandler.HttpErrorHandler
	r := chi.NewRouter()

	r.Route("/notification", func(r chi.Router) {
		r.Route("/", func(r chi.Router) {
			r.Get("/", eh(handler.GetNotifications))
			r.Get("/unread", eh(handler.GetUnreadCount))
		})
		r.Patch("/{notificationID}", eh(handler.MarkAsRead))
		r.Patch("/readAll", eh(handler.MarkAllAsRead))
	})

	return r
}
