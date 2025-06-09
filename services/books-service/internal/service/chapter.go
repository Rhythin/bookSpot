package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/errhandler"
)

func (s *service) AddChapter(ctx context.Context, chapter *entities.Chapter) (err error) {

	// get existingBook by id
	existingBook, err := s.GetBookByID(ctx, chapter.BookID)
	if err != nil {
		return err
	}

	if existingBook == nil {
		return errhandler.NewCustomError(errors.New("book not found"), http.StatusNotFound, "Book not found", false)
	}

	return s.Model.Chapter.AddChapter(ctx, chapter)
}

func (s *service) GetChapterByID(ctx context.Context, bookID string, chapterID string) (*entities.Chapter, error) {

	//get existingBook by id
	existingBook, err := s.GetBookByID(ctx, bookID)
	if err != nil {
		return nil, err
	}

	if existingBook == nil {
		return nil, errhandler.NewCustomError(errors.New("book not found"), http.StatusNotFound, "Book not found", false)
	}

	return s.Model.Chapter.GetChapterByID(ctx, bookID, chapterID)
}

func (s *service) UpdateChapter(ctx context.Context, chapter *entities.Chapter) (err error) {

	// get existingChapter by id
	existingChapter, err := s.GetChapterByID(ctx, chapter.BookID, chapter.ID)
	if err != nil {
		return err
	}

	if existingChapter == nil {
		return errhandler.NewCustomError(errors.New("chapter not found"), http.StatusNotFound, "Chapter not found", false)
	}

	return s.Model.Chapter.UpdateChapter(ctx, chapter)
}

func (s *service) DeleteChapter(ctx context.Context, bookID string, chapterID string) (err error) {

	// get existingChapter by id
	existingChapter, err := s.GetChapterByID(ctx, bookID, chapterID)
	if err != nil {
		return err
	}

	// check if chapter exists
	if existingChapter == nil {
		return errhandler.NewCustomError(errors.New("chapter not found"), http.StatusNotFound, "Chapter not found", false)
	}

	return s.Model.Chapter.DeleteChapter(ctx, bookID, chapterID)
}

func (s *service) GetChapterList(ctx context.Context, req *packets.GetChapterListRequest) (resp *packets.ListChaptersResponse, err error) {

	return s.Model.Chapter.GetChapterList(ctx, req)
}
