package models

type Company struct {
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}