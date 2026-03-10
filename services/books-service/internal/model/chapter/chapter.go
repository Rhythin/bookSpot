package chapter

import (
	"context"
	"net/http"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
	"github.com/rhythin/bookspot/services/shared/tracing"
	"go.opentelemetry.io/otel"
)

func (c *chapter) Add(ctx context.Context, chapter *entities.Chapter) error {
	tr := otel.Tracer("books-model")
	ctx, span := tr.Start(ctx, "AddChapter")
	defer span.End()

	err := c.db.WithContext(ctx).
		Create(chapter).
		Error

	if err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Errorw("failed to add chapter", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to add chapter", false)
	}

	return nil
}

func (c *chapter) GetList(ctx context.Context, req *packets.GetChapterListRequest) (resp *packets.ListChaptersResponse, err error) {
	tr := otel.Tracer("books-model")
	ctx, span := tr.Start(ctx, "GetChapterList")
	defer span.End()

	var chapters []*packets.ChapterDetails
	var totalCount, searchCount int64

	err = c.db.WithContext(ctx).
		Where("book_id = ?", req.BookID).
		Count(&totalCount).
		Where("name LIKE ?", "%"+req.Search+"%").
		Limit(req.Limit).
		Offset(req.Offset).
		Find(&chapters).
		Count(&searchCount).
		Error

	if err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Errorw("failed to get chapter list", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to get chapter list", false)
	}

	return &packets.ListChaptersResponse{
		Chapters:    chapters,
		TotalCount:  totalCount,
		SearchCount: searchCount,
	}, nil
}

func (c *chapter) GetByID(ctx context.Context, bookID string, chapterID string) (*entities.Chapter, error) {
	tr := otel.Tracer("books-model")
	ctx, span := tr.Start(ctx, "GetChapterByID")
	defer span.End()

	var chapter *entities.Chapter

	err := c.db.WithContext(ctx).
		Where("book_id = ?", bookID).
		Where("id = ?", chapterID).
		First(&chapter).
		Error

	if err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Errorw("failed to get chapter by id", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to get chapter by id", false)
	}

	return chapter, nil
}

func (c *chapter) Update(ctx context.Context, chapter *entities.Chapter) error {
	tr := otel.Tracer("books-model")
	ctx, span := tr.Start(ctx, "UpdateChapter")
	defer span.End()

	err := c.db.WithContext(ctx).
		Where("id = ?", chapter.ID).
		Updates(chapter).
		Error

	if err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Errorw("failed to update chapter", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to update chapter", false)
	}

	return nil
}

func (c *chapter) Delete(ctx context.Context, bookID string, chapterID string) error {
	tr := otel.Tracer("books-model")
	ctx, span := tr.Start(ctx, "DeleteChapter")
	defer span.End()

	err := c.db.WithContext(ctx).
		Where("book_id = ?", bookID).
		Where("id = ?", chapterID).
		Delete(&entities.Chapter{}).
		Error

	if err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Errorw("failed to delete chapter", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to delete chapter", false)
	}

	return nil
}

func (c *chapter) GetCount(ctx context.Context, bookIDs []string) (map[string]int64, error) {
	tr := otel.Tracer("books-model")
	ctx, span := tr.Start(ctx, "GetChapterCount")
	defer span.End()

	var chapterCount map[string]int64

	err := c.db.WithContext(ctx).
		Select("book_id, COUNT(*) as count").
		Where("book_id IN (?)", bookIDs).
		Group("book_id").
		Scan(&chapterCount).
		Error

	if err != nil {
		tracing.RecordSpanError(span, err)
		customlogger.S().Errorw("failed to get chapter count", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to get chapter count", false)
	}

	return chapterCount, nil
}
