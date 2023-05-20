package models

type Attendance struct {
	Date  string `json:"date"`
	Items []Item `json:"items"`
}
