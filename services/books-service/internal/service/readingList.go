package service

import (
	"context"

	"github.com/rhythin/bookspot/books-service/internal/entities/packets"
)

func (s *service) AddToReadingList(ctx context.Context, bookID string) (err error) {
	return nil
}

func (s *service) RemoveFromReadingList(ctx context.Context, bookID string) (err error) {
	return nil
}

func (s *service) GetReadingList(ctx context.Context, req *packets.GetReadingListRequest) (*packets.ListBooksResponse, error) {
	resp := &packets.ListBooksResponse{}
	return resp, nil
}

func (s *service) UpdateLastReadChapter(ctx context.Context, bookID string, chapterID string) (err error) {
	return nil
}
