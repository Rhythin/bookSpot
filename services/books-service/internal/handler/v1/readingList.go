package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/custommodel"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"github.com/rhythin/bookspot/services/shared/tracing"
	"go.opentelemetry.io/otel"
)

// AddToReadingList godoc
// @Summary      Add book to reading list
// @Description  Add a book to the user's personal reading list
// @Tags         reading-list
// @Accept       json
// @Produce      json
// @Param        bookID     query     string  true   "Book ID"
// @Param        chapterID  query     string  false  "Starting Chapter ID"
// @Success      200        {object}  map[string]interface{}
// @Failure      400        {object}  map[string]interface{}
// @Router       /reading-list [post]
func (h *handlerV1) AddToReadingList(w http.ResponseWriter, r *http.Request) error {
	tr := otel.Tracer("books-handler")
	ctx, span := tr.Start(r.Context(), "AddToReadingList")
	defer span.End()

	bookID := r.URL.Query().Get("bookID")
	if bookID == "" {
		err := errhandler.NewCustomError(errors.New("book id is required"), http.StatusBadRequest, "Book id is required", false)
		tracing.RecordSpanError(span, err)
		return err
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
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to validate reading list entry", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid reading list entry", false)
	}

	if err := h.Service.AddToReadingList(ctx, readingListEntry); err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}
	return sendResponse(w, nil, http.StatusOK)
}

// RemoveFromReadingList godoc
// @Summary      Remove book from reading list
// @Description  Remove a specific entry from the user's reading list
// @Tags         reading-list
// @Accept       json
// @Produce      json
// @Param        listEntryID  path      string  true  "Reading List Entry ID"
// @Param        bookID       query     string  true  "Book ID"
// @Success      200          {object}  map[string]interface{}
// @Failure      400          {object}  map[string]interface{}
// @Router       /reading-list/{listEntryID} [delete]
func (h *handlerV1) RemoveFromReadingList(w http.ResponseWriter, r *http.Request) error {
	tr := otel.Tracer("books-handler")
	ctx, span := tr.Start(r.Context(), "RemoveFromReadingList")
	defer span.End()

	listEntryID := chi.URLParam(r, "listEntryID")
	if listEntryID == "" {
		err := errhandler.NewCustomError(errors.New("listEntryID is required"), http.StatusBadRequest, "ListEntryID is required", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	bookID := r.URL.Query().Get("bookID")
	if bookID == "" {
		err := errhandler.NewCustomError(errors.New("bookID is required"), http.StatusBadRequest, "BookID is required", false)
		tracing.RecordSpanError(span, err)
		return err
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
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to validate reading list entry", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid reading list entry", false)
	}

	if err := h.Service.RemoveFromReadingList(ctx, readingListEntry); err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}
	return sendResponse(w, nil, http.StatusOK)
}

// UpdateLastReadChapter godoc
// @Summary      Update reading progress
// @Description  Update the last read chapter for a book in the reading list
// @Tags         reading-list
// @Accept       json
// @Produce      json
// @Param        listEntryID  path      string  true  "Reading List Entry ID"
// @Param        bookID       path      string  true  "Book ID"
// @Param        chapterID    path      string  true  "Last Read Chapter ID"
// @Success      200          {object}  map[string]interface{}
// @Failure      400          {object}  map[string]interface{}
// @Router       /reading-list/{listEntryID}/book/{bookID}/chapter/{chapterID} [put]
func (h *handlerV1) UpdateLastReadChapter(w http.ResponseWriter, r *http.Request) error {
	tr := otel.Tracer("books-handler")
	ctx, span := tr.Start(r.Context(), "UpdateLastReadChapter")
	defer span.End()

	listEntryID := chi.URLParam(r, "listEntryID")
	if listEntryID == "" {
		err := errhandler.NewCustomError(errors.New("list entry id is required"), http.StatusBadRequest, "List entry id is required", false)
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("list entry id is required", "Error", err)
		return err
	}

	bookID := chi.URLParam(r, "bookID")
	if bookID == "" {
		err := errhandler.NewCustomError(errors.New("book id is required"), http.StatusBadRequest, "Book id is required", false)
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("book id is required", "Error", err)
		return err
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
		tracing.RecordSpanError(span, err)
		return err
	}
	return sendResponse(w, nil, http.StatusOK)
}

// GetReadingList godoc
// @Summary      Get user's reading list
// @Description  Get a paginated list of books in the user's reading list
// @Tags         reading-list
// @Accept       json
// @Produce      json
// @Param        limit   query     int     false  "Limit"
// @Param        offset  query     int     false  "Offset"
// @Param        search  query     string  false  "Search term"
// @Success      200     {object}  packets.ListReadingListResponse
// @Failure      400     {object}  map[string]interface{}
// @Router       /reading-list [get]
func (h *handlerV1) GetReadingList(w http.ResponseWriter, r *http.Request) error {
	tr := otel.Tracer("books-handler")
	ctx, span := tr.Start(r.Context(), "GetReadingList")
	defer span.End()

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
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to validate request", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	books, err := h.Service.GetReadingList(ctx, req)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}
	return sendResponse(w, books, http.StatusOK)
}
