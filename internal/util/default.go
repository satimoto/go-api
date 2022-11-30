package util

func DefaultBool(val *bool, fallback bool) bool {
	if val != nil {
		return *val
	}

	return fallback
}

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

func DefaultFloat(val *float64, fallback float64) float64 {
	if val != nil {
		return *val
	}

	return fallback
}