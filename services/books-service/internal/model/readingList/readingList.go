package readingList

import (
	"context"

	"github.com/rhythin/bookspot/books-service/internal/entities"
	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
)

func (r *readingList) AddToReadingList(ctx context.Context, bookID string) (err error) {

	return nil
}

func (r *readingList) RemoveFromReadingList(ctx context.Context, bookID string) (err error) {

	return nil
}

func (r *readingList) GetReadingList(ctx context.Context, req *packets.GetReadingListRequest) (entries []*entities.ReadingListEntry, err error) {
	return nil, nil
}

func (r *readingList) UpdateLastReadChapter(ctx context.Context, bookID string, chapterID string) (err error) {
	return nil
}
