package poi

import (
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
)

const (
	RFC3339     = "2006-01-02T15:04:05Z"
	RFC3339Nano = "2006-01-02T15:04:05.999Z"
)

var (
	PARSE_FORMATS = []string{RFC3339Nano, RFC3339}
)

func NewCreatePoiParams(elementDto *ElementDto) db.CreatePoiParams {
	return db.CreatePoiParams{
		Uid:         elementDto.ID,
		Source:      "btcmap",
		TagKey:      "other",
		TagValue:    "other",
		LastUpdated: parseTime(elementDto.UpdatedAt, time.Now()),
	}
}

func parseTime(str string, fallback time.Time) time.Time {
	for _, format := range PARSE_FORMATS {
		if result, err := time.Parse(format, str); err == nil {
			return result
		}
	}

	return fallback
}
