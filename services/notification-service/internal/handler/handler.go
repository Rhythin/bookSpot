package handler

import (
	"github.com/go-playground/validator/v10"
	v1 "github.com/rhythin/bookspot/notification-service/internal/handler/v1"
	"github.com/rhythin/bookspot/notification-service/internal/service"
)

type Handler struct {
	V1      v1.HandlerV1
	Service service.Service
}

func NewHandler(service service.Service, validator *validator.Validate) Handler {
	return Handler{
		V1:      v1.NewHandler(service, validator),
		Service: service,
	}
}
