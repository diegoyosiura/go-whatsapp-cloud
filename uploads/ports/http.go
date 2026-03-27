package ports

import "net/http"

// HTTPDoer wraps the low-level HTTP client allowing test injections.
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}
