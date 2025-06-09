package book

import (
	"context"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"gorm.io/gorm"
)

type Book interface {
	CreateBook(ctx context.Context, book *entities.Book) (err error)
	UpdateBook(ctx context.Context, bookID string, book *entities.Book) (err error)
	DeleteBook(ctx context.Context, bookID string) (err error)
	GetBooks(ctx context.Context, req *packets.GetBooksRequest) (*packets.ListBooksResponse, error)
	GetBookByID(ctx context.Context, bookID string) (*entities.Book, error)
}

type book struct {
	db *gorm.DB
}

func New(db *gorm.DB) Book {
	return &book{
		db: db,
	}
}
