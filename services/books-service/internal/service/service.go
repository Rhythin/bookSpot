package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"github.com/rhythin/bookspot/books-service/internal/model"
)

// Service handles business logic for the books service
type service struct {
	Model     model.Model
	Validator *validator.Validate
}

// NewService creates a new Service instance
func NewService(model model.Model, validator *validator.Validate) Service {
	return &service{
		Model:     model,
		Validator: validator,
	}
}

type Service interface {

	// books
	CreateBook(ctx context.Context, book *entities.Book) error
	GetBookByID(ctx context.Context, bookID string) (*entities.Book, error)
	UpdateBook(ctx context.Context, bookID string, book *entities.Book) error
	DeleteBook(ctx context.Context, bookID string) error
	GetBooks(ctx context.Context, req *packets.GetBooksRequest) (*packets.ListBooksResponse, error)

	// chapters
	AddChapter(ctx context.Context, chapter *entities.Chapter) error
	GetChapterList(ctx context.Context, req *packets.GetChapterListRequest) (*packets.ListChaptersResponse, error)
	GetChapterByID(ctx context.Context, bookID string, chapterID string) (*entities.Chapter, error)
	UpdateChapter(ctx context.Context, chapter *entities.Chapter) error
	DeleteChapter(ctx context.Context, bookID string, chapterID string) error

	// reading list
	AddToReadingList(ctx context.Context, bookID string) error
	RemoveFromReadingList(ctx context.Context, bookID string) error
	GetReadingList(ctx context.Context, req *packets.GetReadingListRequest) (*packets.ListBooksResponse, error)
	UpdateLastReadChapter(ctx context.Context, bookID string, chapterID string) error
}
