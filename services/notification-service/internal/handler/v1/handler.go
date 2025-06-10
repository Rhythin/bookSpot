package v1

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rhythin/bookspot/notification-service/internal/service"
)

type handlerV1 struct {
	Service   service.Service
	Validator *validator.Validate
}

// NewHandler creates a new Handler instance
func NewHandler(service service.Service, validator *validator.Validate) HandlerV1 {
	return &handlerV1{
		Service:   service,
		Validator: validator,
	}
}

type HandlerV1 interface {

	// notifications
	GetNotifications(w http.ResponseWriter, r *http.Request) (err error)
	GetUnreadCount(w http.ResponseWriter, r *http.Request) (err error)
	MarkAsRead(w http.ResponseWriter, r *http.Request) (err error)
	MarkAllAsRead(w http.ResponseWriter, r *http.Request) (err error)
}
