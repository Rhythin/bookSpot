package model

import (
	"github.com/rhythin/bookspot/notification-service/internal/model/notification"
	"gorm.io/gorm"
)

type Model struct {
	Notification notification.Notification
}

func New(db *gorm.DB) Model {
	return Model{
		Notification: notification.New(db),
	}
}
