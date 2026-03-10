package service

import (
	"context"
	"fmt"

	"github.com/rhythin/bookspot/notification-service/internal/entities"
	"github.com/rhythin/bookspot/notification-service/internal/entities/packets"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

func (s *service) GetNotifications(ctx context.Context, userID string) (result []*packets.NotificationDetails, err error) {
	tr := otel.Tracer("notification-service")
	ctx, span := tr.Start(ctx, "GetNotifications")
	defer span.End()

	result, err = s.Model.Notification.GetNotifications(ctx, userID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	return result, err
}

func (s *service) GetUnreadCount(ctx context.Context, userID string) (count int64, err error) {
	tr := otel.Tracer("notification-service")
	ctx, span := tr.Start(ctx, "GetUnreadCount")
	defer span.End()

	count, err = s.Model.Notification.GetUnreadCount(ctx, userID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	return count, err
}

func (s *service) MarkAsRead(ctx context.Context, notificationID string) (err error) {
	tr := otel.Tracer("notification-service")
	ctx, span := tr.Start(ctx, "MarkAsRead")
	defer span.End()

	err = s.Model.Notification.MarkAsRead(ctx, notificationID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	return err
}

func (s *service) MarkAllAsRead(ctx context.Context, userID string) (err error) {
	tr := otel.Tracer("notification-service")
	ctx, span := tr.Start(ctx, "MarkAllAsRead")
	defer span.End()

	err = s.Model.Notification.MarkAllAsRead(ctx, userID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	return err
}

func (s *service) CreateNotification(ctx context.Context, notification *packets.CreateNotificationDetails) (err error) {
	tr := otel.Tracer("notification-service")
	ctx, span := tr.Start(ctx, "CreateNotification")
	defer span.End()

	var notifications []*entities.Notification

	for _, userID := range notification.UserIDs {

		notifications = append(notifications, &entities.Notification{
			UserID:  userID,
			IsRead:  false,
			Title:   fmt.Sprintf("New chapter for %s", notification.BookTitle),
			Message: fmt.Sprintf("Chapter %d: \"%s\" is now available!", notification.ChapterNumber, notification.ChapterTitle),
		})
	}

	err = s.Model.Notification.CreateNotification(ctx, notifications)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	return err
}
