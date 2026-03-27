package ports

import "context"

// ThreadTracker deals with the missing link between the messages SDK and the webhook SDK.
// When Meta accepts a message, it generates a 'wamID'. Webhooks (delivery status, user replies) reference this wamID.
type ThreadTracker interface {
	SaveMessageContext(ctx context.Context, wamID string, contextData map[string]interface{}) error
	ResolveMessageContext(ctx context.Context, wamID string) (map[string]interface{}, error)
}
