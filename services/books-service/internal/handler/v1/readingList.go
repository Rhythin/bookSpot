package v1

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
)

func (h *handlerV1) GetReadingList(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()

	limit := chi.URLParam(r, "limit")
	offset := chi.URLParam(r, "offset")
	search := chi.URLParam(r, "search")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid limit", false)
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid offset", false)
	}

	req := &packets.GetReadingListRequest{
		Limit:  limitInt,
		Offset: offsetInt,
		Search: search,
	}

	if err := h.Validator.Struct(req); err != nil {
		customlogger.S().Warnw("failed to validate request", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	books, err := h.Service.GetReadingList(ctx, req)
	if err != nil {
		return err
	}
	return sendResponse(w, books, http.StatusOK)
}

func (h *handlerV1) AddToReadingList(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()

	bookID := chi.URLParam(r, "book_id")
	if bookID == "" {
		return errhandler.NewCustomError(errors.New("book id is required"), http.StatusBadRequest, "Book id is required", false)
	}

	if err := h.Service.AddToReadingList(ctx, bookID); err != nil {
		return err
	}
	return sendResponse(w, nil, http.StatusOK)
}

func (h *handlerV1) RemoveFromReadingList(w http.ResponseWriter, r *http.Request) error {

	ctx := context.Background()

	bookID := chi.URLParam(r, "book_id")
	if bookID == "" {
		return errhandler.NewCustomError(errors.New("book id is required"), http.StatusBadRequest, "Book id is required", false)
	}

	if err := h.Service.RemoveFromReadingList(ctx, bookID); err != nil {
		return err
	}
	return sendResponse(w, nil, http.StatusOK)
}

func (h *handlerV1) UpdateLastReadChapter(w http.ResponseWriter, r *http.Request) error {

	ctx := context.Background()

	bookID := chi.URLParam(r, "book_id")
	chapterID := chi.URLParam(r, "chapter_id")

	if bookID == "" || chapterID == "" {
		return errhandler.NewCustomError(errors.New("book id and chapter id are required"), http.StatusBadRequest, "Book id and chapter id are required", false)
	}

	if err := h.Service.UpdateLastReadChapter(ctx, bookID, chapterID); err != nil {
		return err
	}
	return sendResponse(w, nil, http.StatusOK)
}
