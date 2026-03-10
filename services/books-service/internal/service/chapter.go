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

func (s *service) AddChapter(ctx context.Context, chapter *entities.Chapter) (err error) {
	tr := otel.Tracer("books-service")
	ctx, span := tr.Start(ctx, "AddChapter")
	defer span.End()

	// get existingBook by id
	existingBook, err := s.GetBookByID(ctx, chapter.BookID)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	if existingBook == nil {
		err := errhandler.NewCustomError(errors.New("book not found"), http.StatusNotFound, "Book not found", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	err = s.Model.Chapter.Add(ctx, chapter)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return err
}

func (s *service) GetChapterByID(ctx context.Context, bookID string, chapterID string) (*entities.Chapter, error) {
	tr := otel.Tracer("books-service")
	ctx, span := tr.Start(ctx, "GetChapterByID")
	defer span.End()

	//get existingBook by id
	existingBook, err := s.GetBookByID(ctx, bookID)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return nil, err
	}

	if existingBook == nil {
		err := errhandler.NewCustomError(errors.New("book not found"), http.StatusNotFound, "Book not found", false)
		tracing.RecordSpanError(span, err)
		return nil, err
	}

	res, err := s.Model.Chapter.GetByID(ctx, bookID, chapterID)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return res, err
}

func (s *service) UpdateChapter(ctx context.Context, chapter *entities.Chapter) (err error) {
	tr := otel.Tracer("books-service")
	ctx, span := tr.Start(ctx, "UpdateChapter")
	defer span.End()

	// get existingChapter by id
	existingChapter, err := s.GetChapterByID(ctx, chapter.BookID, chapter.ID)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	if existingChapter == nil {
		err := errhandler.NewCustomError(errors.New("chapter not found"), http.StatusNotFound, "Chapter not found", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	err = s.Model.Chapter.Update(ctx, chapter)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return err
}

func (s *service) DeleteChapter(ctx context.Context, bookID string, chapterID string) (err error) {
	tr := otel.Tracer("books-service")
	ctx, span := tr.Start(ctx, "DeleteChapter")
	defer span.End()

	// get existingChapter by id
	existingChapter, err := s.GetChapterByID(ctx, bookID, chapterID)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	// check if chapter exists
	if existingChapter == nil {
		err := errhandler.NewCustomError(errors.New("chapter not found"), http.StatusNotFound, "Chapter not found", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	err = s.Model.Chapter.Delete(ctx, bookID, chapterID)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return err
}

func (s *service) GetChapterList(ctx context.Context, req *packets.GetChapterListRequest) (resp *packets.ListChaptersResponse, err error) {
	tr := otel.Tracer("books-service")
	ctx, span := tr.Start(ctx, "GetChapterList")
	defer span.End()

	resp, err = s.Model.Chapter.GetList(ctx, req)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return resp, err
}
