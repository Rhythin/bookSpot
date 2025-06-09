package v1

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/custommodel"
	"github.com/rhythin/bookspot/services/shared/errhandler"
)

func (h *handlerV1) AddToReadingList(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()

	bookID := r.URL.Query().Get("bookID")
	if bookID == "" {
		return errhandler.NewCustomError(errors.New("book id is required"), http.StatusBadRequest, "Book id is required", false)
	}
	chapterID := r.URL.Query().Get("chapterID")

	// TODO: get userID from auth middleware
	userID := r.Context().Value("userID").(string)

	readingListEntry := &entities.ReadingListEntry{
		BookID:          bookID,
		UserID:          userID,
		LastReadChapter: chapterID,
	}

	if err := h.Validator.Struct(readingListEntry); err != nil {
		customlogger.S().Warnw("failed to validate reading list entry", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid reading list entry", false)
	}

	if err := h.Service.AddToReadingList(ctx, readingListEntry); err != nil {
		return err
	}
	return sendResponse(w, nil, http.StatusOK)
}

func (h *handlerV1) RemoveFromReadingList(w http.ResponseWriter, r *http.Request) error {

	ctx := context.Background()

	listEntryID := chi.URLParam(r, "listEntryID")
	if listEntryID == "" {
		return errhandler.NewCustomError(errors.New("listEntryID is required"), http.StatusBadRequest, "ListEntryID is required", false)
	}

	bookID := r.URL.Query().Get("bookID")
	if bookID == "" {
		return errhandler.NewCustomError(errors.New("bookID is required"), http.StatusBadRequest, "BookID is required", false)
	}

	// TODO: get userID from auth middleware
	userID := r.Context().Value("userID").(string)

	readingListEntry := &entities.ReadingListEntry{
		CustomModel: custommodel.CustomModel{
			ID: listEntryID,
		},
		UserID: userID,
		BookID: bookID,
	}

	if err := h.Validator.Struct(readingListEntry); err != nil {
		customlogger.S().Warnw("failed to validate reading list entry", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid reading list entry", false)
	}

	if err := h.Service.RemoveFromReadingList(ctx, readingListEntry); err != nil {
		return err
	}
	return sendResponse(w, nil, http.StatusOK)
}

func (h *handlerV1) UpdateLastReadChapter(w http.ResponseWriter, r *http.Request) error {

	ctx := context.Background()

	listEntryID := chi.URLParam(r, "listEntryID")
	if listEntryID == "" {
		customlogger.S().Warnw("list entry id is required", "Error", errors.New("list entry id is required"))
		return errhandler.NewCustomError(errors.New("list entry id is required"), http.StatusBadRequest, "List entry id is required", false)
	}

	bookID := chi.URLParam(r, "bookID")
	if bookID == "" {
		customlogger.S().Warnw("book id is required", "Error", errors.New("book id is required"))
		return errhandler.NewCustomError(errors.New("book id is required"), http.StatusBadRequest, "Book id is required", false)
	}
	chapterID := chi.URLParam(r, "chapterID")

	// TODO: get userID from auth middleware
	userID := r.Context().Value("userID").(string)

	readingListEntry := &entities.ReadingListEntry{
		CustomModel: custommodel.CustomModel{
			ID: listEntryID,
		},
		BookID:          bookID,
		UserID:          userID,
		LastReadChapter: chapterID,
	}

	if err := h.Service.UpdateLastReadChapter(ctx, readingListEntry); err != nil {
		return err
	}
	return sendResponse(w, nil, http.StatusOK)
}

func (h *handlerV1) GetReadingList(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()

	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	search := r.URL.Query().Get("search")

	// TODO: get userID from auth middleware
	userID := r.Context().Value("userID").(string)

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		customlogger.S().Warnw("failed to convert limit to int", "Error", err)
		customlogger.S().Warnw("using default limit", "Limit", 10)
		limitInt = 10
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		customlogger.S().Warnw("failed to convert offset to int", "Error", err)
		customlogger.S().Warnw("using default offset", "Offset", 0)
	}

	req := &packets.GetReadingListRequest{
		Limit:  limitInt,
		Offset: offsetInt,
		Search: search,
		UserID: userID,
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
