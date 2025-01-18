package books

import (
	"net/http"
)

type transport struct {
	apiKey string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Query().Add("key", t.apiKey)
	return http.DefaultTransport.RoundTrip(req)
}

func NewTransport(apiKey string) *transport {
	return &transport{apiKey: apiKey}
}

func NewClient(apiKey string) *http.Client {
	return &http.Client{Transport: NewTransport(apiKey)}
}
