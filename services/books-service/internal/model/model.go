package model

import (
	"github.com/rhythin/bookspot/books-service/internal/model/book"
	"gorm.io/gorm"
)

// Model contains all the models for the application
type Model struct {
	Book book.Book
}

// New creates a new Model instance
func New(db *gorm.DB) Model {
	return Model{
		Book: book.New(db),
	}
}
