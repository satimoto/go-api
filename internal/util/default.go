package util

func DefaultString(str *string, fallback string) string {
	if str != nil {
		return *str
	}

	return fallback
}

func DefaultInt(val *int, fallback int) int {
	if val != nil {
		return *val
	}

	return fallback
}