package util

import "strings"

func DeleteEmpty (arr []string) []string {
	var result []string
	for _, str := range arr {
		strTrim := strings.TrimSpace(str)

		if strTrim != "" {
			result = append(result, strTrim)
		}
	}

	return result
}