package services

import (
	"context"
	"testing"

	"github.com/diegoyosiura/go-whatsapp-cloud/uploads/domain"
)

type mockUploadClient struct {
	SessionRes *domain.CreateSessionResponse
	DataRes    *domain.UploadDataResponse
	QueryRes   *domain.QueryStatusResponse
	Err        error
}

func (m *mockUploadClient) CreateSession(ctx context.Context, fileLength int, fileType, fileName string) (*domain.CreateSessionResponse, error) {
	return m.SessionRes, m.Err
}

func (m *mockUploadClient) UploadData(ctx context.Context, uploadID string, fileOffset int, fileData []byte, mimeType string) (*domain.UploadDataResponse, error) {
	return m.DataRes, m.Err
}

func (m *mockUploadClient) QueryStatus(ctx context.Context, uploadID string) (*domain.QueryStatusResponse, error) {
	return m.QueryRes, m.Err
}

func TestUploadService_Passthrough(t *testing.T) {
	client := &mockUploadClient{}
	s := NewUploadService(client)

	client.SessionRes = &domain.CreateSessionResponse{ID: "session_id"}
	res, err := s.CreateSession(context.Background(), 1000, "image/jpeg", "test.jpg")
	if err != nil || res.ID != "session_id" {
		t.Errorf("failed CreateSession pass")
	}

	client.DataRes = &domain.UploadDataResponse{Handle: "my_handle"}
	res2, err2 := s.UploadData(context.Background(), "session_id", 0, []byte("ok"), "image/jpeg")
	if err2 != nil || res2.Handle != "my_handle" {
		t.Errorf("failed UploadData pass")
	}

	client.QueryRes = &domain.QueryStatusResponse{FileOffset: 50}
	res3, err3 := s.QueryStatus(context.Background(), "session_id")
	if err3 != nil || res3.FileOffset != 50 {
		t.Errorf("failed QueryStatus pass")
	}
}
