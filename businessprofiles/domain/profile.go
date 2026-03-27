package domain

// ProfileUpdate represents the core information that can be pushed to the
// whatsapp_business_profile edge. It natively mirrors the JSON payload.
type ProfileUpdate struct {
	About          string   `json:"about,omitempty"`
	Address        string   `json:"address,omitempty"`
	Description    string   `json:"description,omitempty"`
	Email          string   `json:"email,omitempty"`
	ProfilePicture string   `json:"profile_picture_handle,omitempty"`
	Vertical       string   `json:"vertical,omitempty"`
	Websites       []string `json:"websites,omitempty"`
}

// ProfileResponse represents what Meta returns when fetching the business profile.
type ProfileResponse struct {
	Data []ProfileData `json:"data"`
}

type ProfileData struct {
	About          string   `json:"about"`
	Address        string   `json:"address"`
	Description    string   `json:"description"`
	Email          string   `json:"email"`
	ProfilePicture string   `json:"profile_picture_url"`
	Websites       []string `json:"websites"`
	Vertical       string   `json:"vertical"`
}
