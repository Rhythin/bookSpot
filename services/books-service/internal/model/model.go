package model

import (
	"github.com/rhythin/bookspot/books-service/internal/model/book"
	"github.com/rhythin/bookspot/books-service/internal/model/chapter"
	"github.com/rhythin/bookspot/books-service/internal/model/readingList"
	"gorm.io/gorm"
)

// Model contains all the models for the application
type Model struct {
	Book        book.Book
	Chapter     chapter.Chapter
	ReadingList readingList.ReadingList
}

// New creates a new Model instance
func New(db *gorm.DB) Model {
	return Model{
		Book:        book.New(db),
		Chapter:     chapter.New(db),
		ReadingList: readingList.New(db),
	}
}
