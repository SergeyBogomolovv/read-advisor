package handler

import (
	"context"
	"errors"
	"log/slog"

	svc "github.com/SergeyBogomolovv/read-advisor/services/books/internal/service"

	pb "github.com/SergeyBogomolovv/read-advisor/lib/api/gen/books"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookService interface {
	Search(ctx context.Context, query string, params svc.ApiParams) (*svc.Volumes, error)
	ByID(ctx context.Context, id string) (*svc.Volume, error)
}

type handler struct {
	logger  *slog.Logger
	service BookService
	pb.UnimplementedBooksServer
}

func NewGRPCHandler(logger *slog.Logger, service BookService) *handler {
	return &handler{
		logger:  logger,
		service: service,
	}
}

func (h *handler) Serve(server *grpc.Server) {
	pb.RegisterBooksServer(server, h)
}

func (h *handler) Search(ctx context.Context, params *pb.SearchParams) (*pb.BookList, error) {
	const op = "handler.Search"
	logger := h.logger.With(slog.String("op", op))

	volumes, err := h.service.Search(ctx, params.Q, svc.ApiParams{
		LangRestrict: params.Lang,
		OrderBy:      params.OrderBy,
		MaxResults:   int(params.MaxResults),
		StartIndex:   int(params.StartIndex),
	})
	if err != nil {
		logger.Error("failed to search books", "err", err)
		return nil, status.Error(codes.Internal, "something went wrong")
	}

	return convertVolumes(volumes), nil
}

func (h *handler) BookByID(ctx context.Context, params *pb.BookID) (*pb.Book, error) {
	const op = "handler.BookByID"
	logger := h.logger.With(slog.String("op", op))

	volume, err := h.service.ByID(context.Background(), params.Id)
	if err != nil {
		if errors.Is(err, svc.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "book not found")
		}
		logger.Error("failed to get book", "err", err)
		return nil, status.Error(codes.Internal, "something went wrong")
	}

	return convertVolume(volume), nil
}
