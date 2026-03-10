package readingList

import (
	"context"
	"errors"
	"net/http"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"github.com/rhythin/bookspot/services/shared/tracing"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

func (r *readingList) Add(ctx context.Context, entry *entities.ReadingListEntry) (err error) {
	tr := otel.Tracer("books-model")
	ctx, span := tr.Start(ctx, "AddToReadingList")
	defer span.End()

	err = r.db.WithContext(ctx).
		Create(entry).
		Error
	if err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to add to reading list", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to add to reading list", false)
	}

	return nil
}

func (r *readingList) Remove(ctx context.Context, entry *entities.ReadingListEntry) (err error) {
	tr := otel.Tracer("books-model")
	ctx, span := tr.Start(ctx, "RemoveFromReadingList")
	defer span.End()

	err = r.db.WithContext(ctx).
		Model(&entities.ReadingListEntry{}).
		Where("id = ?", entry.ID).
		Where("book_id = ?", entry.BookID).
		Delete(entry).
		Error
	if err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to remove from reading list", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to remove from reading list", false)
	}

	return nil
}

func (r *readingList) GetByID(ctx context.Context, entry *entities.ReadingListEntry) (resp *entities.ReadingListEntry, err error) {
	tr := otel.Tracer("books-model")
	ctx, span := tr.Start(ctx, "GetReadingListEntry")
	defer span.End()

	tx := r.db.WithContext(ctx).
		Model(&entities.ReadingListEntry{}).
		Where("id = ?", entry.ID)

	err = tx.First(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to get reading list entry", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to get reading list entry", false)
	}

	return resp, nil
}

func (r *readingList) UpdateLastReadChapter(ctx context.Context, entry *entities.ReadingListEntry) (err error) {
	tr := otel.Tracer("books-model")
	ctx, span := tr.Start(ctx, "UpdateLastReadChapter")
	defer span.End()

	tx := r.db.WithContext(ctx).
		Model(&entities.ReadingListEntry{}).
		Where("id = ?", entry.ID).
		Where("book_id = ?", entry.BookID).
		Where("user_id = ?", entry.UserID)

	err = tx.Update("last_read_chapter", entry.LastReadChapter).
		Error
	if err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to update last read chapter", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to update last read chapter", false)
	}

	return nil
}

func (r *readingList) GetDuplicate(ctx context.Context, entry *entities.ReadingListEntry) (resp *entities.ReadingListEntry, err error) {
	tr := otel.Tracer("books-model")
	ctx, span := tr.Start(ctx, "GetDuplicateReadingListEntry")
	defer span.End()

	tx := r.db.WithContext(ctx).
		Model(&entities.ReadingListEntry{}).
		Where("book_id = ?", entry.BookID).
		Where("user_id = ?", entry.UserID).
		Where("id != ?", entry.ID)

	err = tx.First(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("failed to get reading list entry", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to get reading list entry", false)
	}

	return resp, nil
}

func (r *readingList) GetReadingList(ctx context.Context, req *packets.GetReadingListRequest) (resp *packets.ListReadingListResponse, err error) {
	tr := otel.Tracer("books-model")
	ctx, span := tr.Start(ctx, "GetReadingList")
	defer span.End()

	var entries []*packets.ReadingListEntryDetails
	var totalCount, searchCount int64

	err = r.db.WithContext(ctx).
		Where("user_id = ?", req.UserID).
		Count(&totalCount).
		Where("name LIKE ?", "%"+req.Search+"%").
		Limit(req.Limit).
		Offset(req.Offset).
		Find(&entries).
		Count(&searchCount).
		Error

	if err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Errorw("failed to get reading list", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to get reading list", false)
	}

	return &packets.ListReadingListResponse{
		Entries:     entries,
		TotalCount:  totalCount,
		SearchCount: searchCount,
	}, nil
}
