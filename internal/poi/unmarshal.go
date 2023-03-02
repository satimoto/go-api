package poi

import (
	"encoding/json"
	"io"
)

func UnmarshalDto(body io.ReadCloser) ([]*ElementDto, error) {
	response := []*ElementDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}
