package ports

import (
	"context"
	"net/http"

	"github.com/diegoyosiura/go-whatsapp-cloud/waba/domain"
)

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type WABAClient interface {
	GetAccountInfo(ctx context.Context) (domain.AccountInfo, error)
}
