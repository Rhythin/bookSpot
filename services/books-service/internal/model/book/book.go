package book

import (
	"context"
	"errors"
	"net/http"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"gorm.io/gorm"
)

func (b *book) Create(ctx context.Context, book *entities.Book) (err error) {

	err = b.db.WithContext(ctx).
		Create(book).
		Error

	if err != nil {
		customlogger.S().Errorw("failed to create book", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to create book", false)
	}

	return nil
}

func (b *book) Update(ctx context.Context, bookID string, book *entities.Book) (err error) {

	err = b.db.WithContext(ctx).
		Where("id = ?", bookID).
		Updates(book).
		Error

	if err != nil {
		customlogger.S().Errorw("failed to update book", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to update book", false)
	}

	return nil
}

func (b *book) Delete(ctx context.Context, bookID string) (err error) {

	err = b.db.WithContext(ctx).
		Where("id = ?", bookID).
		Delete(&entities.Book{}).
		Error

	if err != nil {
		customlogger.S().Errorw("failed to delete book", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to delete book", false)
	}

	return nil
}

func (b *book) GetList(ctx context.Context, req *packets.GetBooksRequest) (resp *packets.ListBooksResponse, err error) {
	resp = &packets.ListBooksResponse{}
	var books []*packets.BookDetails

	err = b.db.WithContext(ctx).
		Count(&resp.TotalCount).
		Limit(req.Limit).
		Offset(req.Offset).
		Where("title LIKE ?", "%"+req.Search+"%").
		Find(&books).
		Count(&resp.SearchCount).
		Error

	if err != nil {
		customlogger.S().Errorw("failed to get books", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to get books", false)
	}

	resp = &packets.ListBooksResponse{
		Books:       books,
		TotalCount:  resp.TotalCount,
		SearchCount: resp.SearchCount,
	}

	return resp, nil
}

func (b *book) GetByID(ctx context.Context, bookID string) (*entities.Book, error) {
	var book *entities.Book

	err := b.db.WithContext(ctx).
		Where("id = ?", bookID).
		First(&book).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		customlogger.S().Errorw("failed to get book by id", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to get book by id", false)
	}

	return book, nil
}
