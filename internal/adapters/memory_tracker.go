package adapters

import (
	"context"
	"fmt"
	"sync"
)

// MemoryTracker is a goroutine-safe adapter of ports.ThreadTracker.
// It stores the mapping between Meta Message IDs and arbitrary business contexts.
type MemoryTracker struct {
	store sync.Map
}

// NewMemoryTracker starts a local transient store.
func NewMemoryTracker() *MemoryTracker {
	return &MemoryTracker{}
}

// SaveMessageContext binds a business state object to a specific Meta WAMID.
func (t *MemoryTracker) SaveMessageContext(ctx context.Context, wamID string, contextData map[string]interface{}) error {
	t.store.Store(wamID, contextData)
	return nil
}

// ResolveMessageContext extracts the business state bounded to the returning Webhook payload WAMID.
func (t *MemoryTracker) ResolveMessageContext(ctx context.Context, wamID string) (map[string]interface{}, error) {
	val, exists := t.store.Load(wamID)
	if !exists {
		return nil, fmt.Errorf("no application context bound to message %s", wamID)
	}

	data, ok := val.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("corrupted memory storage schema for message %s", wamID)
	}

	return data, nil
}
