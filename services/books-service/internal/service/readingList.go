package service

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
)

func (s *service) AddToReadingList(ctx context.Context, listEntry *entities.ReadingListEntry) (err error) {
	tr := otel.Tracer("books-service")
	ctx, span := tr.Start(ctx, "AddToReadingList")
	defer span.End()

	// check if entry already exists
	duplicateEntry, err := s.Model.ReadingList.GetDuplicate(ctx, listEntry)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	if duplicateEntry != nil {
		err := errhandler.NewCustomError(errors.New("entry already exists"), http.StatusBadRequest, "Entry already exists", false)
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("entry already exists", "duplicate", duplicateEntry)
		return err
	}

	err = s.Model.ReadingList.Add(ctx, listEntry)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return err
}

func (s *service) RemoveFromReadingList(ctx context.Context, listEntry *entities.ReadingListEntry) (err error) {
	tr := otel.Tracer("books-service")
	ctx, span := tr.Start(ctx, "RemoveFromReadingList")
	defer span.End()

	// check if entry exists
	existingEntry, err := s.Model.ReadingList.GetByID(ctx, listEntry)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	if existingEntry == nil {
		err := errhandler.NewCustomError(errors.New("entry does not exist"), http.StatusBadRequest, "Entry does not exist", false)
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("entry does not exist", "BookID", listEntry.BookID)
		return err
	}

	err = s.Model.ReadingList.Remove(ctx, listEntry)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return err
}

func (s *service) UpdateLastReadChapter(ctx context.Context, listEntry *entities.ReadingListEntry) (err error) {
	tr := otel.Tracer("books-service")
	ctx, span := tr.Start(ctx, "UpdateLastReadChapter")
	defer span.End()

	// check if entry exists
	existingEntry, err := s.Model.ReadingList.GetByID(ctx, listEntry)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return err
	}

	if existingEntry == nil {
		err := errhandler.NewCustomError(errors.New("entry does not exist"), http.StatusBadRequest, "Entry does not exist", false)
		tracing.RecordSpanError(span, err)
		customlogger.S().Warnw("entry does not exist", "BookID", listEntry.BookID)
		return err
	}

	err = s.Model.ReadingList.UpdateLastReadChapter(ctx, listEntry)
	if err != nil {
		tracing.RecordSpanError(span, err)
	}
	return err
}

func (s *service) GetReadingList(ctx context.Context, req *packets.GetReadingListRequest) (*packets.ListReadingListResponse, error) {
	tr := otel.Tracer("books-service")
	ctx, span := tr.Start(ctx, "GetReadingList")
	defer span.End()

	// get reading list
	resp, err := s.Model.ReadingList.GetReadingList(ctx, req)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return nil, err
	}

	bookIDs := make([]string, len(resp.Entries))
	for i, entry := range resp.Entries {
		bookIDs[i] = entry.BookID
	}

	chapterCountMap, err := s.Model.Chapter.GetCount(ctx, bookIDs)
	if err != nil {
		tracing.RecordSpanError(span, err)
		return nil, err
	}

	for i, entry := range resp.Entries {
		resp.Entries[i].ChapterCount = chapterCountMap[entry.BookID]
	}

	return resp, nil
}
