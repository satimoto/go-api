package auth

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	golnurl "github.com/fiatjaf/go-lnurl"
)

func CallbackUrl(version string) string {
	return fmt.Sprintf("%s/%s/lnurl/auth", os.Getenv("API_DOMAIN"), version)
}

func GenerateLnUrl(version, challenge string) (string, error) {
	params := url.Values{}
	params.Add("k1", challenge)
	params.Add("tag", "login")

	generatedUrl := fmt.Sprintf("%s?%s", CallbackUrl(version), params.Encode())
	encodedUrl, err := golnurl.LNURLEncode(generatedUrl)

	if err != nil {
		return encodedUrl, err
	}

	return strings.ToLower(encodedUrl), nil
}
