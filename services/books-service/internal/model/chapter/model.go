package chapter

import (
	"context"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"gorm.io/gorm"
)

type Chapter interface {
	Add(ctx context.Context, chapter *entities.Chapter) error
	GetList(ctx context.Context, req *packets.GetChapterListRequest) (*packets.ListChaptersResponse, error)
	GetByID(ctx context.Context, bookID string, chapterID string) (*entities.Chapter, error)
	Update(ctx context.Context, chapter *entities.Chapter) error
	Delete(ctx context.Context, bookID, chapterID string) error
	GetCount(ctx context.Context, bookIDs []string) (map[string]int64, error)
}

type chapter struct {
	db *gorm.DB
}

func New(db *gorm.DB) Chapter {
	return &chapter{
		db: db,
	}
}
