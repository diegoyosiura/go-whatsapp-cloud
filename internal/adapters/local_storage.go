package adapters

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

// LocalStorage implements the ports.FileStorage interface saving binary streams directly to local disk/OS file system.
type LocalStorage struct {
	basePath string
}

// NewLocalStorage instantiates a file system adapter saving to the provided strict root folder.
func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{
		basePath: basePath,
	}
}

// Save attempts to recursively build directories based on the arbitrary path and flush the bytes array securely.
func (s *LocalStorage) Save(ctx context.Context, path string, data []byte) error {
	fullPath := filepath.Join(s.basePath, path)

	// Build any nonexistent structural tree matching the path
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0750); err != nil {
		return fmt.Errorf("could not create directories to store media: %w", err)
	}

	return os.WriteFile(fullPath, data, 0600)
}

// Get fetches the data slice from the physical disk layer.
func (s *LocalStorage) Get(ctx context.Context, path string) ([]byte, error) {
	fullPath := filepath.Join(s.basePath, path)
	/* #nosec G304 */
	return os.ReadFile(fullPath)
}

// Delete gracefully removes an untempered file string.
func (s *LocalStorage) Delete(ctx context.Context, path string) error {
	fullPath := filepath.Join(s.basePath, path)
	return os.Remove(fullPath)
}
