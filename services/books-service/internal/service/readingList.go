package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/errhandler"
)

func (s *service) AddToReadingList(ctx context.Context, listEntry *entities.ReadingListEntry) (err error) {

	// check if entry already exists
	duplicateEntry, err := s.Model.ReadingList.GetDuplicate(ctx, listEntry)
	if err != nil {
		return err
	}

	if duplicateEntry != nil {
		customlogger.S().Warnw("entry already exists", "duplicate", duplicateEntry)
		return errhandler.NewCustomError(errors.New("entry already exists"), http.StatusBadRequest, "Entry already exists", false)
	}

	return s.Model.ReadingList.Add(ctx, listEntry)
}

func (s *service) RemoveFromReadingList(ctx context.Context, listEntry *entities.ReadingListEntry) (err error) {

	// check if entry exists
	existingEntry, err := s.Model.ReadingList.GetByID(ctx, listEntry)
	if err != nil {
		return err
	}

	if existingEntry == nil {
		customlogger.S().Warnw("entry does not exist", "BookID", listEntry.BookID)
		return errhandler.NewCustomError(errors.New("entry does not exist"), http.StatusBadRequest, "Entry does not exist", false)
	}

	return s.Model.ReadingList.Remove(ctx, listEntry)
}

func (s *service) UpdateLastReadChapter(ctx context.Context, listEntry *entities.ReadingListEntry) (err error) {

	// check if entry exists
	existingEntry, err := s.Model.ReadingList.GetByID(ctx, listEntry)
	if err != nil {
		return err
	}

	if existingEntry == nil {
		customlogger.S().Warnw("entry does not exist", "BookID", listEntry.BookID)
		return errhandler.NewCustomError(errors.New("entry does not exist"), http.StatusBadRequest, "Entry does not exist", false)
	}

	return s.Model.ReadingList.UpdateLastReadChapter(ctx, listEntry)
}

func (s *service) GetReadingList(ctx context.Context, req *packets.GetReadingListRequest) (*packets.ListReadingListResponse, error) {

	// get reading list
	resp, err := s.Model.ReadingList.GetReadingList(ctx, req)
	if err != nil {
		return nil, err
	}

	bookIDs := make([]string, len(resp.Entries))
	for i, entry := range resp.Entries {
		bookIDs[i] = entry.BookID
	}

	chapterCountMap, err := s.Model.Chapter.GetChapterCount(ctx, bookIDs)
	if err != nil {
		return nil, err
	}

	for i, entry := range resp.Entries {
		resp.Entries[i].ChapterCount = chapterCountMap[entry.BookID]
	}

	return resp, nil
}
