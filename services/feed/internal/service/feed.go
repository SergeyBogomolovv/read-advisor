package service

import (
	"log/slog"

	pb "github.com/SergeyBogomolovv/read-advisor/lib/api/gen/books"
)

type FeedStorage interface{}

type feedService struct {
	logger  *slog.Logger
	books   pb.BooksClient
	storage FeedStorage
}

func NewFeedService(logger *slog.Logger, books pb.BooksClient, storage FeedStorage) *feedService {
	return &feedService{
		logger:  logger,
		books:   books,
		storage: storage,
	}
}
