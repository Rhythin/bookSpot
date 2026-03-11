package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"github.com/rhythin/bookspot/services/shared/tracing"
	"go.opentelemetry.io/otel"
)

// CreateBook godoc
// @Summary      Create a new book
// @Description  Create a new book with the provided details
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      entities.Book  true  "Book Details"
// @Success      201   {object}  entities.Book
// @Failure      400   {object}  map[string]interface{}
// @Router       /books [post]
func (h *handlerV1) CreateBook(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("books-handler")
	ctx, span := tr.Start(r.Context(), "CreateBook")
	defer span.End()

	if r.Body == http.NoBody {
		err := errhandler.NewCustomError(errors.New("no body provided"), http.StatusBadRequest, "No body provided", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	var book entities.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to decode book", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid body", false)
	}

	// validate book
	if err := h.Validator.Struct(book); err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to validate book", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid book", false)
	}

	if err := h.Service.CreateBook(ctx, &book); err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	return sendResponse(w, book, http.StatusCreated)
}

// GetBookByID godoc
// @Summary      Get book by ID
// @Description  Get detailed information about a specific book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book_id  path      string  true  "Book ID"
// @Success      200      {object}  entities.Book
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Router       /books/{book_id} [get]
func (h *handlerV1) GetBookByID(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("books-handler")
	ctx, span := tr.Start(r.Context(), "GetBookByID")
	defer span.End()

	bookID := chi.URLParam(r, "book_id")

	if bookID == "" {
		err := errhandler.NewCustomError(errors.New("book id is required"), http.StatusBadRequest, "Book id is required", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	book, err := h.Service.GetBookByID(ctx, bookID)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	return sendResponse(w, book, http.StatusOK)
}

// UpdateBook godoc
// @Summary      Update book
// @Description  Update details of an existing book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book_id  path      string         true  "Book ID"
// @Param        book     body      entities.Book  true  "Book Update Details"
// @Success      200      {object}  entities.Book
// @Failure      400      {object}  map[string]interface{}
// @Router       /books/{book_id} [put]
func (h *handlerV1) UpdateBook(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("books-handler")
	ctx, span := tr.Start(r.Context(), "UpdateBook")
	defer span.End()

	bookID := chi.URLParam(r, "book_id")

	if bookID == "" {
		err := errhandler.NewCustomError(errors.New("book id is required"), http.StatusBadRequest, "Book id is required", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	var book entities.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to decode book", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid body", false)
	}

	// validate book
	if err := h.Validator.Struct(book); err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to validate book", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid book", false)
	}

	if err := h.Service.UpdateBook(ctx, bookID, &book); err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	return sendResponse(w, book, http.StatusOK)
}

// DeleteBook godoc
// @Summary      Delete book
// @Description  Permanently delete a book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book_id  path      string  true  "Book ID"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]interface{}
// @Router       /books/{book_id} [delete]
func (h *handlerV1) DeleteBook(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("books-handler")
	ctx, span := tr.Start(r.Context(), "DeleteBook")
	defer span.End()

	bookID := chi.URLParam(r, "book_id")

	if bookID == "" {
		err := errhandler.NewCustomError(errors.New("book id is required"), http.StatusBadRequest, "Book id is required", false)
		tracing.RecordSpanError(span, err)
		return err
	}

	if err := h.Service.DeleteBook(ctx, bookID); err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	return sendResponse(w, nil, http.StatusOK)
}

// GetBooks godoc
// @Summary      List books
// @Description  Get a paginated list of books
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        limit   query     int     false  "Limit"
// @Param        offset  query     int     false  "Offset"
// @Param        search  query     string  false  "Search term"
// @Success      200     {object}  packets.ListBooksResponse
// @Failure      500     {object}  map[string]interface{}
// @Router       /books [get]
func (h *handlerV1) GetBooks(w http.ResponseWriter, r *http.Request) (err error) {
	tr := otel.Tracer("books-handler")
	ctx, span := tr.Start(r.Context(), "GetBooks")
	defer span.End()

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

	req := &packets.GetBooksRequest{
		Limit:  limitInt,
		Offset: offsetInt,
		Search: search,
	}

	if err := h.Validator.Struct(req); err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to validate request", "Error", err)
		return errhandler.NewCustomError(err, http.StatusBadRequest, "Invalid request", false)
	}

	books, err := h.Service.GetBooks(ctx, req)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}
	return sendResponse(w, books, http.StatusOK)
}
