package service

import (
	"context"
	"log/slog"

	pb "github.com/SergeyBogomolovv/read-advisor/lib/api/gen/books"
	d "github.com/SergeyBogomolovv/read-advisor/services/feed/internal/domain"
	"golang.org/x/sync/errgroup"
)

type PrefStorage interface {
	SavePreference(ctx context.Context, userID int64, pref d.Preference, value string, priority int) error
}

type prefService struct {
	logger  *slog.Logger
	books   pb.BooksClient
	storage PrefStorage
}

func NewPreferenceService(logger *slog.Logger, books pb.BooksClient, storage PrefStorage) *prefService {
	return &prefService{
		logger:  logger,
		books:   books,
		storage: storage,
	}
}

func (s *prefService) AddPreference(ctx context.Context, userID int64, bookID string, prefType d.PreferenceType) error {
	const op = "preference.AddPreference"
	logger := s.logger.With(slog.String("op", op), slog.String("book_id", bookID))

	book, err := s.books.BookByID(ctx, &pb.BookID{Id: bookID})
	if err != nil {
		logger.Error("failed to get book", "err", err)
		return err
	}

	var priority int
	switch prefType {
	case d.PreferenceTypeDefault:
		priority = 3
	case d.PreferenceTypeSaved:
		priority = 2
	case d.PreferenceTypeLiked:
		priority = 1
	}

	var eg errgroup.Group
	for _, author := range book.Authors {
		eg.Go(func() error {
			return s.storage.SavePreference(ctx, userID, d.PreferenceAuthor, author, priority)
		})
	}

	for _, category := range book.Categories {
		eg.Go(func() error {
			return s.storage.SavePreference(ctx, userID, d.PreferenceCategory, category, priority)
		})
	}

	if book.Publisher != "" {
		eg.Go(func() error {
			return s.storage.SavePreference(ctx, userID, d.PreferencePublisher, book.Publisher, priority)
		})
	}

	return eg.Wait()
}
