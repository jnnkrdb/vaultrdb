package helpers

import (
	"encoding/json"
	"io"
)

func ReadObjectFromJSONBody(dst interface{}, body io.ReadCloser) error {
	if err := json.NewDecoder(body).Decode(dst); err != nil {
		return err
	}
	return nil
}
