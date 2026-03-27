package services

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/waba/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/waba/ports"
)

// wabaService sits between the facade and the HTTP networking ports.
type wabaService struct {
	client ports.WABAClient
}

// NewWABAService injects the abstract API contract.
func NewWABAService(client ports.WABAClient) ports.WABAService {
	return &wabaService{
		client: client,
	}
}

// GetAccountInfo simply invokes the specific endpoint for the global WABA scope.
func (s *wabaService) GetAccountInfo(ctx context.Context) (domain.AccountInfo, error) {
	return s.client.GetAccountInfo(ctx)
}
