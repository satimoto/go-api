package util

import (
	"database/sql"

	"github.com/satimoto/go-datastore/geom"
)

func NullGeometry(value geom.NullGeometry4326) (*geom.Geometry4326, error) {
	if value.Valid {
		return &value.Geometry4326, nil
	}

	return nil, nil
}

func NullInt(value sql.NullInt32) (*int, error) {
	if value.Valid {
		val := int(value.Int32)
		return &val, nil
	}

	return nil, nil
}

func NullString(value sql.NullString) (*string, error) {
	if value.Valid {
		return &value.String, nil
	}

	return nil, nil
}

func NullTime(value sql.NullTime, layout string) (*string, error) {
	if value.Valid {
		val := value.Time.Format(layout)
		return &val, nil
	}

	return nil, nil
}
