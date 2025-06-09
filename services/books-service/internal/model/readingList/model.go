package readingList

import (
	"context"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"gorm.io/gorm"
)

type ReadingList interface {
	Add(ctx context.Context, entry *entities.ReadingListEntry) error
	Remove(ctx context.Context, entry *entities.ReadingListEntry) error
	GetByID(ctx context.Context, entry *entities.ReadingListEntry) (*entities.ReadingListEntry, error)
	GetDuplicate(ctx context.Context, entry *entities.ReadingListEntry) (*entities.ReadingListEntry, error)
	UpdateLastReadChapter(ctx context.Context, entry *entities.ReadingListEntry) error
	GetReadingList(ctx context.Context, req *packets.GetReadingListRequest) (*packets.ListReadingListResponse, error)
}

type readingList struct {
	db *gorm.DB
}

func New(db *gorm.DB) ReadingList {
	return &readingList{
		db: db,
	}
}
