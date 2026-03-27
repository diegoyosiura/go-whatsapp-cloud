package services

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/diegoyosiura/go-whatsapp-cloud/media/domain"
)

type mockMediaClient struct {
	Info     domain.MediaInfo
	InfoErr  error
	BinData  io.ReadCloser
	BinErr   error
	UploadID string
	UpErr    error
}

func (m *mockMediaClient) GetMediaInfo(ctx context.Context, mediaID string) (domain.MediaInfo, error) {
	return m.Info, m.InfoErr
}

func (m *mockMediaClient) DownloadBinary(ctx context.Context, mediaURL string) (io.ReadCloser, error) {
	return m.BinData, m.BinErr
}

func (m *mockMediaClient) UploadBinary(ctx context.Context, fileReader io.Reader, mimeType string) (string, error) {
	return m.UploadID, m.UpErr
}

func TestMediaService_DownloadMedia(t *testing.T) {
	mockClient := &mockMediaClient{}
	service := NewMediaService(mockClient)

	t.Run("Successful Download Pipeline", func(t *testing.T) {
		mockClient.Info = domain.MediaInfo{URL: "https://meta.com/secure_url"}
		mockClient.BinData = io.NopCloser(strings.NewReader("fake_binary_data"))
		mockClient.InfoErr = nil
		mockClient.BinErr = nil

		reader, err := service.DownloadMedia(context.Background(), "media_id_123")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		defer reader.Close()
		data, _ := io.ReadAll(reader)
		if string(data) != "fake_binary_data" {
			t.Errorf("expected binary data 'fake_binary_data', got %s", string(data))
		}
	})

	t.Run("Fails at GetMediaInfo", func(t *testing.T) {
		mockClient.InfoErr = errors.New("api limit reached")

		_, err := service.DownloadMedia(context.Background(), "media_id_123")
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("Fails at DownloadBinary", func(t *testing.T) {
		mockClient.InfoErr = nil // Success at info
		mockClient.Info = domain.MediaInfo{URL: "https://meta.com/secure_url"}
		mockClient.BinErr = errors.New("network timeout") // Fails at binary fetch

		_, err := service.DownloadMedia(context.Background(), "media_id_123")
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}

func TestMediaService_UploadMedia(t *testing.T) {
	mockClient := &mockMediaClient{}
	service := NewMediaService(mockClient)

	t.Run("Successful Upload", func(t *testing.T) {
		mockClient.UploadID = "wamid.success.123"
		mockClient.UpErr = nil

		id, err := service.UploadMedia(context.Background(), strings.NewReader("bin"), "image/jpeg")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if id != "wamid.success.123" {
			t.Errorf("expected id 'wamid.success.123', got %s", id)
		}
	})
}
