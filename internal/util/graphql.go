package util

import (
	"database/sql"

	"github.com/satimoto/go-datastore/pkg/geom"
)

func NullGeometry(value geom.NullGeometry4326) (*geom.Geometry4326, error) {
	if value.Valid {
		return &value.Geometry4326, nil
	}

	return nil, nil
}

func NullFloat(value sql.NullFloat64) (*float64, error) {
	if value.Valid {
		return &value.Float64, nil
	}

	return nil, nil
}

func NullInt(i interface{}) (*int, error) {
	switch t := i.(type) {
	case sql.NullInt32:
		if t.Valid {
			val := int(t.Int32)
			return &val, nil
		}
	case sql.NullInt64:
		if t.Valid {
			val := int(t.Int64)
			return &val, nil
		}
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
