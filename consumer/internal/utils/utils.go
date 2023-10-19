package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// UnmarshalInto deserializes the content in `in` into an uninitialized struct.
// Returns nil if the operation succeeded, otherwise any error encountered.
func UnmarshalInto[T any](out *T, in io.Reader) error {
	var unmarshallErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(in)
	decoder.DisallowUnknownFields() // Error on unknown fields
	if err := decoder.Decode(out); err != nil {
		if errors.As(err, &unmarshallErr) {
			return fmt.Errorf("incorrect type in field: %v", unmarshallErr.Field)
		} else {
			return fmt.Errorf("bad request: %v", err)
		}
	}
	return nil
}
