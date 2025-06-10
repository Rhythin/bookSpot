package service

import (
	"context"
	"fmt"

	"github.com/rhythin/bookspot/notification-service/internal/entities"
	"github.com/rhythin/bookspot/notification-service/internal/entities/packets"
)

func (s *service) GetNotifications(ctx context.Context, userID string) (result []*packets.NotificationDetails, err error) {
	return s.Model.Notification.GetNotifications(ctx, userID)
}

func (s *service) GetUnreadCount(ctx context.Context, userID string) (count int64, err error) {
	return s.Model.Notification.GetUnreadCount(ctx, userID)
}

func (s *service) MarkAsRead(ctx context.Context, notificationID string) (err error) {
	return s.Model.Notification.MarkAsRead(ctx, notificationID)
}

func (s *service) MarkAllAsRead(ctx context.Context, userID string) (err error) {
	return s.Model.Notification.MarkAllAsRead(ctx, userID)
}

func (s *service) CreateNotification(ctx context.Context, notification *packets.CreateNotificationDetails) (err error) {

	var notifications []*entities.Notification

	for _, userID := range notification.UserIDs {

		notifications = append(notifications, &entities.Notification{
			UserID:  userID,
			IsRead:  false,
			Title:   fmt.Sprintf("New chapter for %s", notification.BookTitle),
			Message: fmt.Sprintf("Chapter %d: \"%s\" is now available!", notification.ChapterNumber, notification.ChapterTitle),
		})
	}

	return s.Model.Notification.CreateNotification(ctx, notifications)
}
