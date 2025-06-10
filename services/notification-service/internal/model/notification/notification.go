package notification

import (
	"context"

	"github.com/rhythin/bookspot/notification-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
)

func (n *notification) GetNotifications(ctx context.Context, userID string) (result []*packets.NotificationDetails, err error) {

	err = n.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&result).Error
	if err != nil {
		customlogger.S().Warnw("failed to get notifications for user", "userID", userID, "error", err)
		return nil, err
	}

	return result, nil
}

func (n *notification) GetUnreadCount(ctx context.Context, userID string) (count int64, err error) {

	err = n.db.WithContext(ctx).
		Where("user_id = ? AND is_read = ?", userID, false).
		Find(&count).
		Error
	if err != nil {
		customlogger.S().Warnw("failed to get unread count for user", "userID", userID, "error", err)
		return 0, err
	}

	return count, nil
}

func (n *notification) MarkAsRead(ctx context.Context, notificationID string) (err error) {

	err = n.db.WithContext(ctx).
		Where("id = ?", notificationID).
		Update("is_read", true).
		Error
	if err != nil {
		customlogger.S().Warnw("failed to mark notification as read", "notificationID", notificationID, "error", err)
		return err
	}

	return nil
}

func (n *notification) MarkAllAsRead(ctx context.Context, userID string) (err error) {

	err = n.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Update("is_read", true).
		Error
	if err != nil {
		customlogger.S().Warnw("failed to mark all notifications as read for user", "userID", userID, "error", err)
		return err
	}

	return nil
}
