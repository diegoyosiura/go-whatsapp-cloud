package ports

import "net/http"

// HTTPDoer wraps net/http to enable intercepting HTTP calls during unit testing.
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}
