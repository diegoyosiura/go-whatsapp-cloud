package adapters

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestLocalStorage_E2ECycle(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "whatsapp_sdk_storage_test_*")
	if err != nil {
		t.Fatalf("failed creating mock dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	storage := NewLocalStorage(tmpDir)
	filePath := "mediastream/file.txt"

	t.Run("Saves successfully in tree", func(t *testing.T) {
		err := storage.Save(context.Background(), filePath, []byte("Hello Meta SDK"))
		if err != nil {
			t.Fatalf("unexpected save error: %v", err)
		}

		fullPath := filepath.Join(tmpDir, filePath)
		actual, err := os.ReadFile(fullPath)
		if err != nil {
			t.Fatalf("did not create file %s", fullPath)
		}

		if string(actual) != "Hello Meta SDK" {
			t.Errorf("content mismatched: %s", string(actual))
		}
	})

	t.Run("Retrieves successfully", func(t *testing.T) {
		data, err := storage.Get(context.Background(), filePath)
		if err != nil {
			t.Fatalf("unexpected get error: %v", err)
		}

		if string(data) != "Hello Meta SDK" {
			t.Errorf("get content mismatch")
		}
	})

	t.Run("Deletes cleanly", func(t *testing.T) {
		err := storage.Delete(context.Background(), filePath)
		if err != nil {
			t.Fatalf("delete error: %v", err)
		}

		_, err = storage.Get(context.Background(), filePath)
		if err == nil {
			t.Errorf("file was allegedly deleted, but Get recovered it")
		}
	})

	t.Run("Get fails gracefully on nonexistent", func(t *testing.T) {
		_, err := storage.Get(context.Background(), "bogus/not_there.bin")
		if err == nil {
			t.Error("expected missing file err")
		}
	})
}
