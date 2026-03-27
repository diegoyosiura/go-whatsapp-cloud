package domain

// CreateSessionResponse holds the Upload ID handle to be used in consequent chunk actions.
type CreateSessionResponse struct {
	ID string `json:"id"`
}

// UploadDataResponse retrieves the final handle required by the API destination (e.g., Profile Picture).
type UploadDataResponse struct {
	Handle string `json:"h"`
}

// QueryStatusResponse reports the cursor chunk successfully committed into Meta storage.
type QueryStatusResponse struct {
	ID         string `json:"id"`
	FileOffset int    `json:"file_offset"`
}
