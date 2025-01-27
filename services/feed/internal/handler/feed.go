package handler

import (
	"context"
	"log/slog"

	pb "github.com/SergeyBogomolovv/read-advisor/lib/api/gen/feed"
	"google.golang.org/grpc"
)

type FeedService interface{}

type grpcHandler struct {
	logger  *slog.Logger
	service FeedService
	pb.UnimplementedFeedServer
}

func NewFeedHandler(logger *slog.Logger, service FeedService) *grpcHandler {
	return &grpcHandler{
		logger:  logger,
		service: service,
	}
}

func (h *grpcHandler) Register(server *grpc.Server) {
	pb.RegisterFeedServer(server, h)
}

func (h *grpcHandler) ForUser(ctx context.Context, in *pb.UserID) (*pb.Book, error) {
	return nil, nil
}
