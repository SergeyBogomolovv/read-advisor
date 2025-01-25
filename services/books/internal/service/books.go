package service

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/SergeyBogomolovv/read-advisor/services/books/internal/lib/transport"
	"github.com/SergeyBogomolovv/read-advisor/services/books/pkg/e"
	"github.com/SergeyBogomolovv/read-advisor/services/books/pkg/utils"
)

const basePath = "https://www.googleapis.com/books/v1/volumes"

type service struct {
	client *http.Client
	logger *slog.Logger
}

func NewBooksService(logger *slog.Logger, apiKey string) *service {
	return &service{
		client: &http.Client{Transport: transport.NewApiTransport(apiKey)},
		logger: logger,
	}
}

func (s *service) Search(ctx context.Context, query string, params ApiParams) (*Volumes, error) {
	const op = "service.SearchBooks"
	logger := s.logger.With(slog.String("op", op))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, basePath, nil)
	if err != nil {
		logger.Error("failed to create request", "err", err)
		return nil, err
	}
	values := params.ToValues()
	values.Set("q", query)
	values.Set("fields", volumesFields)
	req.URL.RawQuery = values.Encode()

	resp, err := s.doRequest(logger, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	volumes := new(Volumes)
	if err := utils.DecodeBody(resp.Body, volumes); err != nil {
		logger.Error("failed to decode response", "err", err)
		return nil, err
	}
	return volumes, nil
}

func (s *service) ByID(ctx context.Context, id string) (*Volume, error) {
	const op = "service.ByID"
	logger := s.logger.With(slog.String("op", op), slog.String("id", id))

	url := fmt.Sprintf("%s/%s?fields=%s", basePath, id, volumeFields)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		logger.Error("failed to create request", "err", err)
		return nil, err
	}

	resp, err := s.doRequest(logger, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	volume := new(Volume)
	if err := utils.DecodeBody(resp.Body, volume); err != nil {
		logger.Error("failed to decode response", "err", err)
		return nil, err
	}
	return volume, nil
}

func (s *service) doRequest(logger *slog.Logger, req *http.Request) (*http.Response, error) {
	resp, err := s.client.Do(req)
	if err != nil {
		logger.Error("failed to do request", "err", err)
		return nil, e.Wrap("failed to do request", err)
	}

	if resp.StatusCode != http.StatusOK {
		logger.Error("failed to request books", "status", resp.StatusCode)
		return nil, fmt.Errorf("failed to request books: %d", resp.StatusCode)
	}

	return resp, nil
}
