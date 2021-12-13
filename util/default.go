package util

func DefaultString(str *string, def string) string {
	if str != nil {
		return *str
	}

	return def
}