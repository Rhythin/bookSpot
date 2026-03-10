package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"github.com/rhythin/bookspot/services/shared/tracing"
	"go.opentelemetry.io/otel"
)

func (s *service) CreateBook(ctx context.Context, book *entities.Book) (err error) {
	tr := otel.Tracer("books-service")
	ctx, span := tr.Start(ctx, "CreateBook")
	defer span.End()

	err = s.Model.Book.Create(ctx, book)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return err
}

func (s *service) GetBookByID(ctx context.Context, bookID string) (*entities.Book, error) {
	tr := otel.Tracer("books-service")
	ctx, span := tr.Start(ctx, "GetBookByID")
	defer span.End()

	res, err := s.Model.Book.GetByID(ctx, bookID)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return res, err
}

func (s *service) UpdateBook(ctx context.Context, bookID string, book *entities.Book) (err error) {
	tr := otel.Tracer("books-service")
	ctx, span := tr.Start(ctx, "UpdateBook")
	defer span.End()

	// get existingBook by id
	existingBook, err := s.GetBookByID(ctx, bookID)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	if existingBook == nil {
		err := errhandler.NewCustomError(errors.New("book not found"), http.StatusNotFound, "Book not found", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	err = s.Model.Book.Update(ctx, bookID, book)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return err
}

func (s *service) DeleteBook(ctx context.Context, bookID string) (err error) {
	tr := otel.Tracer("books-service")
	ctx, span := tr.Start(ctx, "DeleteBook")
	defer span.End()

	// get existingBook by id
	existingBook, err := s.GetBookByID(ctx, bookID)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	if existingBook == nil {
		err := errhandler.NewCustomError(errors.New("book not found"), http.StatusNotFound, "Book not found", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	err = s.Model.Book.Delete(ctx, bookID)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return err
}

func (s *service) GetBooks(ctx context.Context, req *packets.GetBooksRequest) (resp *packets.ListBooksResponse, err error) {
	tr := otel.Tracer("books-service")
	ctx, span := tr.Start(ctx, "GetBooks")
	defer span.End()

	resp, err = s.Model.Book.GetList(ctx, req)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return nil, err
	}

	// create an array of book ids
	bookIDs := make([]string, len(resp.Books))
	for i, book := range resp.Books {
		bookIDs[i] = book.ID
	}

	// get chapter count for each book
	chapterCountMap, err := s.Model.Chapter.GetCount(ctx, bookIDs)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return nil, err
	}

	// set chapter count for each book
	for _, book := range resp.Books {
		book.ChapterCount = chapterCountMap[book.ID]
	}

	return resp, nil
}
