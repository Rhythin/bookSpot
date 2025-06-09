package chapter

import (
	"context"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"gorm.io/gorm"
)

type Chapter interface {
	AddChapter(ctx context.Context, chapter *entities.Chapter) error
	GetChapterList(ctx context.Context, req *packets.GetChapterListRequest) (*packets.ListChaptersResponse, error)
	GetChapterByID(ctx context.Context, bookID string, chapterID string) (*entities.Chapter, error)
	UpdateChapter(ctx context.Context, chapter *entities.Chapter) error
	DeleteChapter(ctx context.Context, bookID, chapterID string) error
	GetChapterCount(ctx context.Context, bookIDs []string) (map[string]int64, error)
}

type chapter struct {
	db *gorm.DB
}

func New(db *gorm.DB) Chapter {
	return &chapter{
		db: db,
	}
}
