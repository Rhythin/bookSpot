package handler

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/go-playground/validator/v10"
	"github.com/rhythin/bookspot/notification-service/internal/service"
)

type eventHandler struct {
	service   service.Service
	validator *validator.Validate
}

func NewEventHandler(service service.Service, validator *validator.Validate) EventHandler {
	return &eventHandler{service: service, validator: validator}
}

type EventHandler interface {
	SendNotification(ctx context.Context, headers map[string]string, message *sarama.ConsumerMessage) error
}
