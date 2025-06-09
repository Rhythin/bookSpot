package chapter

import (
	"context"
	"net/http"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
)

func (c *chapter) AddChapter(ctx context.Context, chapter *entities.Chapter) error {

	err := c.db.WithContext(ctx).
		Create(chapter).
		Error

	if err != nil {
		customlogger.S().Errorw("failed to add chapter", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to add chapter", false)
	}

	return nil
}

func (c *chapter) GetChapterList(ctx context.Context, req *packets.GetChapterListRequest) ([]*entities.Chapter, error) {
	var chapters []*entities.Chapter

	err := c.db.WithContext(ctx).
		Limit(req.Limit).
		Offset(req.Offset).
		Where("book_id = ?", req.BookID).
		Where("name LIKE ?", "%"+req.Search+"%").
		Find(&chapters).
		Error

	if err != nil {
		customlogger.S().Errorw("failed to get chapter list", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to get chapter list", false)
	}

	return chapters, nil
}

func (c *chapter) GetChapterByID(ctx context.Context, bookID string, chapterID string) (*entities.Chapter, error) {
	var chapter *entities.Chapter

	err := c.db.WithContext(ctx).
		Where("book_id = ?", bookID).
		Where("id = ?", chapterID).
		First(&chapter).
		Error

	if err != nil {
		customlogger.S().Errorw("failed to get chapter by id", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to get chapter by id", false)
	}

	return chapter, nil
}

func (c *chapter) UpdateChapter(ctx context.Context, chapter *entities.Chapter) error {

	err := c.db.WithContext(ctx).
		Where("id = ?", chapter.ID).
		Updates(chapter).
		Error

	if err != nil {
		customlogger.S().Errorw("failed to update chapter", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to update chapter", false)
	}

	return nil
}

func (c *chapter) DeleteChapter(ctx context.Context, bookID string, chapterID string) error {

	err := c.db.WithContext(ctx).
		Where("book_id = ?", bookID).
		Where("id = ?", chapterID).
		Delete(&entities.Chapter{}).
		Error

	if err != nil {
		customlogger.S().Errorw("failed to delete chapter", "Error", err)
		return errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to delete chapter", false)
	}

	return nil
}

func (c *chapter) GetChapterCount(ctx context.Context, bookIDs []string) (map[string]int, error) {
	var chapterCount map[string]int

	err := c.db.WithContext(ctx).
		Select("book_id, COUNT(*) as count").
		Where("book_id IN (?)", bookIDs).
		Group("book_id").
		Scan(&chapterCount).
		Error

	if err != nil {
		customlogger.S().Errorw("failed to get chapter count", "Error", err)
		return nil, errhandler.NewCustomError(err, http.StatusInternalServerError, "Failed to get chapter count", false)
	}

	return chapterCount, nil
}
