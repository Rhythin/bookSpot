package entities

import "github.com/rhythin/bookspot/services/shared/custommodel"

type Notification struct {
	custommodel.CustomModel
	UserID  string `json:"userID"`
	Title   string `json:"title"`
	Message string `json:"message"`
	IsRead  bool   `json:"isRead"`
}
