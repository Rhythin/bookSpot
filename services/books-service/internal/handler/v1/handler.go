package v1

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rhythin/bookspot/books-service/internal/service"
)

type handlerV1 struct {
	Service   service.Service
	Validator *validator.Validate
}

// NewHandler creates a new Handler instance
func NewHandler(service service.Service, validator *validator.Validate) HandlerV1 {
	return &handlerV1{
		Service:   service,
		Validator: validator,
	}
}

type HandlerV1 interface {

	// books
	GetBooks(w http.ResponseWriter, r *http.Request) (err error) // query params: limit, offset, search
	GetBookByID(w http.ResponseWriter, r *http.Request) (err error)
	CreateBook(w http.ResponseWriter, r *http.Request) (err error)
	UpdateBook(w http.ResponseWriter, r *http.Request) (err error)
	DeleteBook(w http.ResponseWriter, r *http.Request) (err error)

	// chapters
	AddChapter(w http.ResponseWriter, r *http.Request) (err error)
	GetChapterList(w http.ResponseWriter, r *http.Request) (err error)
	GetChapterByID(w http.ResponseWriter, r *http.Request) (err error)
	UpdateChapter(w http.ResponseWriter, r *http.Request) (err error)
	DeleteChapter(w http.ResponseWriter, r *http.Request) (err error)

	// reading list
	GetReadingList(w http.ResponseWriter, r *http.Request) (err error)
	AddToReadingList(w http.ResponseWriter, r *http.Request) (err error)
	RemoveFromReadingList(w http.ResponseWriter, r *http.Request) (err error)
	UpdateLastReadChapter(w http.ResponseWriter, r *http.Request) (err error)
}
