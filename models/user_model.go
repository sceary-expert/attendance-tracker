package models

type User struct {
	Name             string       `json:"name,omitempty" validate:"required"`
	EmailId          string       `json:"email-id,omitempty" validate:"required" `
	Password         string       `json:"password,omitempty" validate:"required"`
	SavedAttendances []Attendance `json:"saved-attendances"`
}
