package v1

import (
	"github.com/go-chi/chi/v5"
	v1 "github.com/rhythin/bookspot/books-service/internal/handler/v1"
	errhandler "github.com/rhythin/bookspot/services/shared/errhandler"
)

func NewRouter(handler v1.HandlerV1) chi.Router {
	eh := errhandler.HttpErrorHandler
	r := chi.NewRouter()

	r.Route("/book", func(r chi.Router) {
		r.Post("/", eh(handler.CreateBook))
		r.Get("/", eh(handler.GetBooks)) // query params: limit, offset, search
		r.Get("/{bookID}", eh(handler.GetBookByID))
		r.Put("/{bookID}", eh(handler.UpdateBook))
		r.Delete("/{bookID}", eh(handler.DeleteBook))

		// Nested chapters route under book
		r.Route("/{bookID}/chapters", func(r chi.Router) {
			r.Post("/", eh(handler.AddChapter))
			r.Get("/", eh(handler.GetChapterList)) // query params: limit, offset, search
			r.Get("/{chapterID}", eh(handler.GetChapterByID))
			r.Put("/{chapterID}", eh(handler.UpdateChapter))
			r.Delete("/{chapterID}", eh(handler.DeleteChapter))
		})
	})

	r.Route("/readingList", func(r chi.Router) {
		r.Post("/", eh(handler.AddToReadingList)) // url params: bookID, chapterID
		r.Delete("/{listEntryID}", eh(handler.RemoveFromReadingList))
		r.Get("/", eh(handler.GetReadingList))                       // query params: limit, offset, search
		r.Patch("/{listEntryID}", eh(handler.UpdateLastReadChapter)) // url params: listEntryID
	})

	return r
}
