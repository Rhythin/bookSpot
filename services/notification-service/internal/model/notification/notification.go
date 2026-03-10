package notification

import (
	"context"

	"github.com/rhythin/bookspot/notification-service/internal/entities"
	"github.com/rhythin/bookspot/notification-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

func (n *notification) GetNotifications(ctx context.Context, userID string) (result []*packets.NotificationDetails, err error) {
	tr := otel.Tracer("notification-model")
	ctx, span := tr.Start(ctx, "GetNotifications")
	defer span.End()

	err = n.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&result).Error
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		customlogger.S().Warnw("failed to get notifications for user", "userID", userID, "error", err)
		return nil, err
	}

	return result, nil
}

func (n *notification) GetUnreadCount(ctx context.Context, userID string) (count int64, err error) {
	tr := otel.Tracer("notification-model")
	ctx, span := tr.Start(ctx, "GetUnreadCount")
	defer span.End()

	err = n.db.WithContext(ctx).
		Where("user_id = ? AND is_read = ?", userID, false).
		Find(&count).
		Error
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		customlogger.S().Warnw("failed to get unread count for user", "userID", userID, "error", err)
		return 0, err
	}

	return count, nil
}

func (n *notification) MarkAsRead(ctx context.Context, notificationID string) (err error) {
	tr := otel.Tracer("notification-model")
	ctx, span := tr.Start(ctx, "MarkAsRead")
	defer span.End()

	err = n.db.WithContext(ctx).
		Where("id = ?", notificationID).
		Update("is_read", true).
		Error
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		customlogger.S().Warnw("failed to mark notification as read", "notificationID", notificationID, "error", err)
		return err
	}

	return nil
}

func (n *notification) MarkAllAsRead(ctx context.Context, userID string) (err error) {
	tr := otel.Tracer("notification-model")
	ctx, span := tr.Start(ctx, "MarkAllAsRead")
	defer span.End()

	err = n.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Update("is_read", true).
		Error
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		customlogger.S().Warnw("failed to mark all notifications as read for user", "userID", userID, "error", err)
		return err
	}

	return nil
}

func (n *notification) CreateNotification(ctx context.Context, notifications []*entities.Notification) (err error) {
	tr := otel.Tracer("notification-model")
	ctx, span := tr.Start(ctx, "CreateNotification")
	defer span.End()

	err = n.db.WithContext(ctx).
		Create(&notifications).Error
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		customlogger.S().Warnw("failed to create notification", "error", err)
		return err
	}

	return nil
}
