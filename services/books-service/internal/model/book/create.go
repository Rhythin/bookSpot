package book

import (
	"context"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"go.uber.org/zap"
)

func (b *book) CreateBook(ctx context.Context, book *entities.Book) (err error) {

	err = b.db.WithContext(ctx).
		Create(book).
		Error

	if err != nil {
		zap.S().Errorw("failed to create book", "Error", err)
		return err
	}

	return nil
}
