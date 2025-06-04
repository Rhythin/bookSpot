package entities

import "github.com/rhythin/bookspot/services/shared/custommodel"

type ReadingListEntry struct {
	custommodel.CustomModel
	UserID          string `gorm:"type:uuid;not null;index"`
	MangaID         string `gorm:"type:uuid;not null;index"`
	LastReadChapter int    `gorm:"default:0"` // e.g., 0 means not started
}
