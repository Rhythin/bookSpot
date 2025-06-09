package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/errhandler"
)

func (s *service) CreateBook(ctx context.Context, book *entities.Book) (err error) {

	return s.Model.Book.CreateBook(ctx, book)
}

func (s *service) GetBookByID(ctx context.Context, bookID string) (*entities.Book, error) {
	return s.Model.Book.GetBookByID(ctx, bookID)
}

func (s *service) UpdateBook(ctx context.Context, bookID string, book *entities.Book) (err error) {

	// get existingBook by id
	existingBook, err := s.GetBookByID(ctx, bookID)
	if err != nil {
		return err
	}

	if existingBook == nil {
		return errhandler.NewCustomError(errors.New("book not found"), http.StatusNotFound, "Book not found", false)
	}

	return s.Model.Book.UpdateBook(ctx, bookID, book)
}

func (s *service) DeleteBook(ctx context.Context, bookID string) (err error) {

	// get existingBook by id
	existingBook, err := s.GetBookByID(ctx, bookID)
	if err != nil {
		return err
	}

	if existingBook == nil {
		return errhandler.NewCustomError(errors.New("book not found"), http.StatusNotFound, "Book not found", false)
	}

	return s.Model.Book.DeleteBook(ctx, bookID)
}

func (s *service) GetBooks(ctx context.Context, req *packets.GetBooksRequest) (resp *packets.ListBooksResponse, err error) {

	resp, err = s.Model.Book.GetBooks(ctx, req)
	if err != nil {
		return nil, err
	}

	// create an array of book ids
	bookIDs := make([]string, len(resp.Books))
	for i, book := range resp.Books {
		bookIDs[i] = book.ID
	}

	// get chapter count for each book
	chapterCountMap, err := s.Model.Chapter.GetChapterCount(ctx, bookIDs)
	if err != nil {
		return nil, err
	}

	// set chapter count for each book
	for _, book := range resp.Books {
		book.ChapterCount = chapterCountMap[book.ID]
	}

	return resp, nil
}
