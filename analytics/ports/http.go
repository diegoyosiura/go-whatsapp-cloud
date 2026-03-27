package ports

import "net/http"

// HTTPDoer wraps the core execution of network requests to facilitate Unit Testing (TDD). 
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}
