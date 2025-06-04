package book

import (
	"context"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"gorm.io/gorm"
)

type Book interface {
	CreateBook(ctx context.Context, book *entities.Book) (err error)
}

type book struct {
	db *gorm.DB
}

func New(db *gorm.DB) Book {
	return &book{
		db: db,
	}
}
