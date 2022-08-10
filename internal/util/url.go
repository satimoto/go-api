package util

import (
	"fmt"
	"strings"
)

func AddURLSlash(url string) string {
	if strings.IndexAny(url, "/") == 0 {
		return url
	}

	return fmt.Sprintf("/%s", url)
}

func URLLocale(url, locale, defaultLocale string) string {
	if locale == defaultLocale {
		return url
	}

	return fmt.Sprintf("/%s%s", locale, AddURLSlash(url))
}