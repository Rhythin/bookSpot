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

func (h *handlerV1) CreateBook(w http.ResponseWriter, r *http.Request) (err error) {
	ctx := context.Background()

	if r.Body == http.NoBody {
		return errhandler.NewCustomError(errors.New("no body provided"), http.StatusBadRequest, "No body provided", false)
	}

	var book entities.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		customlogger.S().Warnw("failed to decode book", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid body", false)
	}

	// validate book
	if err := h.Validator.Struct(book); err != nil {
		customlogger.S().Warnw("failed to validate book", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid book", false)
	}

	if err := h.Service.CreateBook(ctx, &book); err != nil {
		return err
	}

	return sendResponse(w, book, http.StatusCreated)
}

func (h *handlerV1) GetBookByID(w http.ResponseWriter, r *http.Request) (err error) {
	ctx := context.Background()

	bookID := chi.URLParam(r, "book_id")

	if bookID == "" {
		return errhandler.NewCustomError(errors.New("book id is required"), http.StatusBadRequest, "Book id is required", false)
	}

	book, err := h.Service.GetBookByID(ctx, bookID)
	if err != nil {
		return err
	}

	return sendResponse(w, book, http.StatusOK)
}

func (h *handlerV1) UpdateBook(w http.ResponseWriter, r *http.Request) (err error) {
	ctx := context.Background()

	bookID := chi.URLParam(r, "book_id")

	if bookID == "" {
		return errhandler.NewCustomError(errors.New("book id is required"), http.StatusBadRequest, "Book id is required", false)
	}

	var book entities.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		customlogger.S().Warnw("failed to decode book", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid body", false)
	}

	// validate book
	if err := h.Validator.Struct(book); err != nil {
		customlogger.S().Warnw("failed to validate book", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid book", false)
	}

	if err := h.Service.UpdateBook(ctx, bookID, &book); err != nil {
		return err
	}

	return sendResponse(w, book, http.StatusOK)
}

func (h *handlerV1) DeleteBook(w http.ResponseWriter, r *http.Request) (err error) {
	ctx := context.Background()

	bookID := chi.URLParam(r, "book_id")

	if bookID == "" {
		return errhandler.NewCustomError(errors.New("book id is required"), http.StatusBadRequest, "Book id is required", false)
	}

	if err := h.Service.DeleteBook(ctx, bookID); err != nil {
		return err
	}

	return sendResponse(w, nil, http.StatusOK)
}

func (h *handlerV1) GetBooks(w http.ResponseWriter, r *http.Request) (err error) {
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

	req := &packets.GetBooksRequest{
		Limit:  limitInt,
		Offset: offsetInt,
		Search: search,
	}

	if err := h.Validator.Struct(req); err != nil {
		customlogger.S().Warnw("failed to validate request", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	books, err := h.Service.GetBooks(ctx, req)
	if err != nil {
		return err
	}
	return sendResponse(w, books, http.StatusOK)
}
