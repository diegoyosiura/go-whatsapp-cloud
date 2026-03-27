package domain

// MediaInfo represents the response given by WhatsApp Graph API when requesting Media Details by ID.
type MediaInfo struct {
	URL      string `json:"url"`
	MimeType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
	FileSize int    `json:"file_size"`
	ID       string `json:"id"`
	// Error fields if Meta responds with failure
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// UploadResponse represents the response containing the ID after a media is correctly uploaded.
type UploadResponse struct {
	ID string `json:"h"`
}
