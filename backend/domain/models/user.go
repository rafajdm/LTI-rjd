package models

type User struct {
	ID              int               `json:"id"`
	Name            string            `json:"name"`
	Email           string            `json:"email"`
	ContactInfo     map[string]string `json:"contact_info"` // JSONB field
	ResumeLink      string            `json:"resume_link"`
	CurrentPosition string            `json:"current_position"`
	Location        string            `json:"location"`
}
