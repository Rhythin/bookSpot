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
		r.Get("/{book_id}", eh(handler.GetBookByID))
		r.Put("/{book_id}", eh(handler.UpdateBook))
		r.Delete("/{book_id}", eh(handler.DeleteBook))

		// Nested chapters route under book
		r.Route("/{book_id}/chapters", func(r chi.Router) {
			r.Post("/", eh(handler.AddChapter))
			r.Get("/", eh(handler.GetChapterList)) // query params: limit, offset, search
			r.Get("/{chapter_id}", eh(handler.GetChapterByID))
			r.Put("/{chapter_id}", eh(handler.UpdateChapter))
			r.Delete("/{chapter_id}", eh(handler.DeleteChapter))
		})
	})

	r.Route("/reading-list", func(r chi.Router) {
		r.Post("/{book_id}", eh(handler.AddToReadingList))
		r.Delete("/{book_id}", eh(handler.RemoveFromReadingList))
		r.Get("/", eh(handler.GetReadingList)) // query params: limit, offset, search
		r.Patch("/{book_id}", eh(handler.UpdateLastReadChapter))
	})

	return r
}
