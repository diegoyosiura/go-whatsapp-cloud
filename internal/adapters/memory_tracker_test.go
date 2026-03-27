package adapters

import (
	"context"
	"testing"
)

func TestMemoryTracker(t *testing.T) {
	tracker := NewMemoryTracker()

	t.Run("Saves and Resolves contexts correctly", func(t *testing.T) {
		ctxData := map[string]interface{}{
			"user_id":   "8855",
			"flow":      "onboarding_step_2",
		}

		err := tracker.SaveMessageContext(context.Background(), "wamid.12345", ctxData)
		if err != nil {
			t.Fatalf("unexpected save error")
		}

		resolved, err := tracker.ResolveMessageContext(context.Background(), "wamid.12345")
		if err != nil {
			t.Fatalf("unexpected resolve error")
		}

		if resolved["user_id"] != "8855" || resolved["flow"] != "onboarding_step_2" {
			t.Errorf("Context drifted or vanished in memory")
		}
	})

	t.Run("Resolving unfetched WAMID returns error", func(t *testing.T) {
		_, err := tracker.ResolveMessageContext(context.Background(), "wamid.ghost")
		if err == nil {
			t.Fatalf("expected error catching nonexistent message tracking")
		}
	})
}
