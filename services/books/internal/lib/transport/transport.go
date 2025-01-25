package transport

import (
	"net/http"
)

type transport struct {
	apiKey string
}

func NewApiTransport(apiKey string) http.RoundTripper {
	return &transport{apiKey: apiKey}
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Query().Set("key", t.apiKey)
	return http.DefaultTransport.RoundTrip(req)
}
