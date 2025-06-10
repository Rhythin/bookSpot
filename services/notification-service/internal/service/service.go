package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/rhythin/bookspot/notification-service/internal/entities/packets"
	"github.com/rhythin/bookspot/notification-service/internal/model"
)

// Service handles business logic for the notification service
type service struct {
	Model     model.Model
	Validator *validator.Validate
}

// NewService creates a new Service instance
func NewService(model model.Model, validator *validator.Validate) Service {
	return &service{
		Model:     model,
		Validator: validator,
	}
}

type Service interface {

	// notifications
	GetNotifications(ctx context.Context, userID string) ([]*packets.NotificationDetails, error)
	GetUnreadCount(ctx context.Context, userID string) (int64, error)
	MarkAsRead(ctx context.Context, notificationID string) error
	MarkAllAsRead(ctx context.Context, userID string) error

	// kafka
	CreateNotification(ctx context.Context, notification *packets.CreateNotificationDetails) error
}
