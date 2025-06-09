package entities

import "github.com/rhythin/bookspot/services/shared/custommodel"

type ReadingListEntry struct {
	custommodel.CustomModel
	UserID          string `gorm:"type:uuid;not null;uniqueIndex:idx_user_book" validate:"required"`
	BookID          string `gorm:"type:uuid;not null;uniqueIndex:idx_user_book" validate:"required"`
	LastReadChapter string `gorm:"default:''"` // e.g., "" means not started
}
