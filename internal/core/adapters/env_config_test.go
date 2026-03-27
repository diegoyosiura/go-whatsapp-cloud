package adapters_test

import (
	"context"
	"os"
	"testing"

	"github.com/diegoyosiura/go-whatsapp-cloud/internal/core/adapters"
)

func TestEnvConfigReader_GetString(t *testing.T) {
	os.Setenv("STRING_TEST_KEY", "string_value")
	defer os.Unsetenv("STRING_TEST_KEY")

	reader := adapters.NewEnvConfigReader()
	ctx := context.Background()

	tests := []struct {
		name     string
		key      string
		expected string
	}{
		{
			name:     "existing string key",
			key:      "STRING_TEST_KEY",
			expected: "string_value",
		},
		{
			name:     "non-existing key returns empty",
			key:      "MISSING_KEY",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reader.GetString(ctx, tt.key)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.expected {
				t.Errorf("GetString() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestEnvConfigReader_GetInt(t *testing.T) {
	os.Setenv("INT_TEST_KEY", "123")
	os.Setenv("INVALID_INT_KEY", "abc")
	defer func() {
		os.Unsetenv("INT_TEST_KEY")
		os.Unsetenv("INVALID_INT_KEY")
	}()

	reader := adapters.NewEnvConfigReader()
	ctx := context.Background()

	tests := []struct {
		name        string
		key         string
		expected    int
		expectError bool
	}{
		{
			name:        "existing int key",
			key:         "INT_TEST_KEY",
			expected:    123,
			expectError: false,
		},
		{
			name:        "invalid int key",
			key:         "INVALID_INT_KEY",
			expected:    0,
			expectError: true,
		},
		{
			name:        "missing key returns err",
			key:         "MISSING_KEY",
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reader.GetInt(ctx, tt.key)
			if (err != nil) != tt.expectError {
				t.Errorf("GetInt() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if got != tt.expected {
				t.Errorf("GetInt() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestEnvConfigReader_GetBool(t *testing.T) {
	os.Setenv("BOOL_TEST_KEY", "true")
	os.Setenv("INVALID_BOOL_KEY", "not-a-bool")
	defer func() {
		os.Unsetenv("BOOL_TEST_KEY")
		os.Unsetenv("INVALID_BOOL_KEY")
	}()

	reader := adapters.NewEnvConfigReader()
	ctx := context.Background()

	tests := []struct {
		name        string
		key         string
		expected    bool
		expectError bool
	}{
		{
			name:        "existing bool key",
			key:         "BOOL_TEST_KEY",
			expected:    true,
			expectError: false,
		},
		{
			name:        "invalid bool key",
			key:         "INVALID_BOOL_KEY",
			expected:    false,
			expectError: true,
		},
		{
			name:        "missing key returns err",
			key:         "MISSING_KEY",
			expected:    false,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reader.GetBool(ctx, tt.key)
			if (err != nil) != tt.expectError {
				t.Errorf("GetBool() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if got != tt.expected {
				t.Errorf("GetBool() = %v, want %v", got, tt.expected)
			}
		})
	}
}
