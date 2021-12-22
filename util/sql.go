package util

import (
	"database/sql"
)

func SqlNullInt32(i interface{}) sql.NullInt32 {
	n := sql.NullInt32{}

	switch t := i.(type) {
	case *int32:
		if t == nil {
			n.Scan(nil)
		} else {
			n.Scan(*t)
		}
	default:
		n.Scan(t)
	}

	return n
}

func SqlNullInt64(i interface{}) sql.NullInt64 {
	n := sql.NullInt64{}

	switch t := i.(type) {
	case *int64:
		if t == nil {
			n.Scan(nil)
		} else {
			n.Scan(*t)
		}
	default:
		n.Scan(t)
	}

	return n
}

func SqlNullString(i interface{}) sql.NullString {
	n := sql.NullString{}

	switch t := i.(type) {
	case *string:
		if t == nil {
			n.Scan(nil)
		} else {
			n.Scan(*t)
		}
	default:
		n.Scan(t)
	}
	
	return n
}
