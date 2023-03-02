package transportation

import (
	"net/http"
)

type HTTPRequester interface {
	Do(req *http.Request) (*http.Response, error)
}
