package v1

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
)

func (h *handlerV1) AddChapter(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()

	bookID := chi.URLParam(r, "book_id")

	if bookID == "" {
		return errhandler.NewCustomError(errors.New("book id is required"), http.StatusBadRequest, "Book id is required", false)
	}

	if r.Body == http.NoBody {
		return errhandler.NewCustomError(errors.New("no body provided"), http.StatusBadRequest, "No body provided", false)
	}

	var chapter entities.Chapter
	if err := json.NewDecoder(r.Body).Decode(&chapter); err != nil {
		customlogger.S().Warnw("failed to decode chapter", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid body", false)
	}

	chapter.BookID = bookID

	// validate chapter
	if err := h.Validator.Struct(chapter); err != nil {
		customlogger.S().Warnw("failed to validate chapter", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid chapter", false)
	}

	if err := h.Service.AddChapter(ctx, &chapter); err != nil {
		return err
	}

	return sendResponse(w, chapter, http.StatusCreated)
}

func (h *handlerV1) GetChapterList(w http.ResponseWriter, r *http.Request) error {

	ctx := context.Background()

	bookID := chi.URLParam(r, "book_id")
	if bookID == "" {
		return errhandler.NewCustomError(errors.New("book id is required"), http.StatusBadRequest, "Book id is required", false)
	}

	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	search := r.URL.Query().Get("search")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		customlogger.S().Warnw("failed to convert limit to int", "Error", err)
		customlogger.S().Info("using default limit", "Limit", 10)
		limitInt = 10
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		customlogger.S().Warnw("failed to convert offset to int", "Error", err)
		customlogger.S().Info("using default offset", "Offset", 0)
	}

	req := &packets.GetChapterListRequest{
		BookID: bookID,
		Limit:  limitInt,
		Offset: offsetInt,
		Search: search,
	}

	if err := h.Validator.Struct(req); err != nil {
		customlogger.S().Warnw("failed to validate request", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	chapters, err := h.Service.GetChapterList(ctx, req)
	if err != nil {
		return err
	}
	return sendResponse(w, chapters, http.StatusOK)
}

func (h *handlerV1) GetChapterByID(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()

	bookID := chi.URLParam(r, "book_id")
	chapterID := chi.URLParam(r, "chapter_id")

	if bookID == "" || chapterID == "" {
		return errhandler.NewCustomError(errors.New("book id and chapter id are required"), http.StatusBadRequest, "Book id and chapter id are required", false)
	}

	chapter, err := h.Service.GetChapterByID(ctx, bookID, chapterID)
	if err != nil {
		return err
	}
	return sendResponse(w, chapter, http.StatusOK)
}

func (h *handlerV1) UpdateChapter(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()

	bookID := chi.URLParam(r, "book_id")
	chapterID := chi.URLParam(r, "chapter_id")

	if bookID == "" || chapterID == "" {
		return errhandler.NewCustomError(errors.New("book id and chapter id are required"), http.StatusBadRequest, "Book id and chapter id are required", false)
	}

	if r.Body == http.NoBody {
		return errhandler.NewCustomError(errors.New("no body provided"), http.StatusBadRequest, "No body provided", false)
	}

	var chapter entities.Chapter
	if err := json.NewDecoder(r.Body).Decode(&chapter); err != nil {
		customlogger.S().Warnw("failed to decode chapter", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid body", false)
	}

	// validate chapter
	if err := h.Validator.Struct(chapter); err != nil {
		customlogger.S().Warnw("failed to validate chapter", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid chapter", false)
	}

	chapter.BookID = bookID
	chapter.ID = chapterID

	if err := h.Service.UpdateChapter(ctx, &chapter); err != nil {
		return err
	}
	return sendResponse(w, chapter, http.StatusOK)
}

func (h *handlerV1) DeleteChapter(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()

	bookID := chi.URLParam(r, "book_id")
	chapterID := chi.URLParam(r, "chapter_id")

	if bookID == "" || chapterID == "" {
		return errhandler.NewCustomError(errors.New("book id and chapter id are required"), http.StatusBadRequest, "Book id and chapter id are required", false)
	}

	if err := h.Service.DeleteChapter(ctx, bookID, chapterID); err != nil {
		return err
	}
	return sendResponse(w, nil, http.StatusOK)
}
