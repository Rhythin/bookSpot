package packets

import "time"

type NotificationDetails struct {
	ID        string    `json:"id"`
	BookID    string    `json:"bookID"`
	ChapterID string    `json:"chapterID"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"isRead"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
}

type NotificationsResponse struct {
	Notifications []*NotificationDetails `json:"notifications"`
}
