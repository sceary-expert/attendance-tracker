package models

type Item struct {
	Name      string `json:"name"`
	IsPresent bool   `json:"is-present"`
}
