package readingList

import (
	"context"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"gorm.io/gorm"
)

type ReadingList interface {
	AddToReadingList(ctx context.Context, bookID string) error
	RemoveFromReadingList(ctx context.Context, bookID string) error
	GetReadingList(ctx context.Context, req *packets.GetReadingListRequest) ([]*entities.ReadingListEntry, error)
	UpdateLastReadChapter(ctx context.Context, bookID string, chapterID string) error
}

type readingList struct {
	db *gorm.DB
}

func New(db *gorm.DB) ReadingList {
	return &readingList{
		db: db,
	}
}
