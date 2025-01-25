package handler_test

import (
	"context"
	"testing"

	pb "github.com/SergeyBogomolovv/read-advisor/lib/api/gen/books"
	"github.com/SergeyBogomolovv/read-advisor/lib/common/test"
	"github.com/SergeyBogomolovv/read-advisor/services/books/internal/handler"
	svc "github.com/SergeyBogomolovv/read-advisor/services/books/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGRPCHandler_Search(t *testing.T) {
	ctx := context.Background()
	bs := new(mockBookService)
	h := handler.NewGRPCHandler(test.NewTestLogger(), bs)

	t.Run("success", func(t *testing.T) {
		bs.On("Search", ctx, "test", mock.Anything).Return(&svc.Volumes{TotalItems: 10}, nil)
		res, err := h.Search(ctx, &pb.SearchParams{Q: "test"})
		assert.NoError(t, err)
		assert.Equal(t, int32(10), res.Total)
		bs.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		bs.On("Search", ctx, "error", mock.Anything).Return(nil, assert.AnError)
		_, err := h.Search(ctx, &pb.SearchParams{Q: "error"})
		assert.Error(t, err)
		bs.AssertExpectations(t)
	})
}

func TestGRPCHandler_BookByID(t *testing.T) {
	ctx := context.Background()
	bs := new(mockBookService)
	h := handler.NewGRPCHandler(test.NewTestLogger(), bs)

	t.Run("success", func(t *testing.T) {
		bs.On("ByID", ctx, "test").Return(&svc.Volume{Id: "test"}, nil)
		res, err := h.BookByID(ctx, &pb.BookID{Id: "test"})
		assert.NoError(t, err)
		assert.Equal(t, res.Id, "test")
		bs.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		bs.On("ByID", ctx, "invalid").Return(nil, svc.ErrNotFound)
		_, err := h.BookByID(ctx, &pb.BookID{Id: "invalid"})
		assert.Error(t, err)
		bs.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		bs.On("ByID", ctx, "err").Return(nil, assert.AnError)
		_, err := h.BookByID(ctx, &pb.BookID{Id: "err"})
		assert.Error(t, err)
		bs.AssertExpectations(t)
	})
}

type mockBookService struct {
	mock.Mock
}

func (m *mockBookService) Search(ctx context.Context, q string, params svc.ApiParams) (*svc.Volumes, error) {
	args := m.Called(ctx, q, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*svc.Volumes), nil
}

func (m *mockBookService) ByID(ctx context.Context, id string) (*svc.Volume, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*svc.Volume), nil
}
