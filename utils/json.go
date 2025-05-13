package utils

import (
	"encoding/json"
	"io"
)

// EncodeJSON encodes given object into JSON writer
func EncodeJSON[T any](w io.Writer, v T) error {
	return json.NewEncoder(w).Encode(v)
}

// DecodeJSON decodes given JSON reader into given object
func DecodeJSON[T any](r io.Reader, v *T) error {
	return json.NewDecoder(r).Decode(v)
}
