package packets

import "time"

type NotificationDetails struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"isRead"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
}

type NotificationsResponse struct {
	Notifications []*NotificationDetails `json:"notifications"`
}
