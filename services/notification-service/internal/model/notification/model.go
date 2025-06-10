package notification

import (
	"context"

	"github.com/rhythin/bookspot/notification-service/internal/entities"
	"github.com/rhythin/bookspot/notification-service/internal/entities/packets"
	"gorm.io/gorm"
)

type notification struct {
	db *gorm.DB
}

func New(db *gorm.DB) Notification {
	return &notification{db: db}
}

type Notification interface {
	GetNotifications(ctx context.Context, userID string) ([]*packets.NotificationDetails, error)
	GetUnreadCount(ctx context.Context, userID string) (int64, error)
	MarkAsRead(ctx context.Context, notificationID string) error
	MarkAllAsRead(ctx context.Context, userID string) error
	CreateNotification(ctx context.Context, notifications []*entities.Notification) error
}
