package utils

import (
	"encoding/json"
	"io"
)

func DecodeBody(data io.Reader, v any) error {
	return json.NewDecoder(data).Decode(v)
}
