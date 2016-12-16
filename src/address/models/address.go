package models

type Address struct {
	PostCode string `json:"PLZ"`
	Street   string
	City     string
}
