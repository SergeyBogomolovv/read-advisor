package books

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

const volumesPath = "https://www.googleapis.com/books/v1/volumes"

type service struct {
	client *http.Client
	logger *slog.Logger
}

func New(logger *slog.Logger, apiKey string) *service {
	return &service{
		client: NewClient(apiKey),
		logger: logger,
	}
}

func (s *service) Volumes(ctx context.Context, params ApiParams) (*Volumes, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, volumesPath, nil)
	req.URL.RawQuery = params.String()
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	volumes := new(Volumes)
	if err := json.NewDecoder(resp.Body).Decode(volumes); err != nil {
		return nil, err
	}
	return volumes, nil
}
