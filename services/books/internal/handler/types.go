package handler

import (
	"context"

	svc "github.com/SergeyBogomolovv/read-advisor/services/books/internal/service"
)

type BookService interface {
	Search(ctx context.Context, query string, params svc.ApiParams) (*svc.Volumes, error)
	ByID(ctx context.Context, id string) (*svc.Volume, error)
}
