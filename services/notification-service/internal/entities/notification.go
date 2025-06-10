package entities

import "github.com/rhythin/bookspot/services/shared/custommodel"

type Notification struct {
	custommodel.CustomModel
	UserID    string
	BookID    string
	ChapterID string
	Message   string
	IsRead    bool
}
