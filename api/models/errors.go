package models

type Errors []Error

type Error struct {
	Parameter string `json:"parameter"`
	Issue     string `json:"error"`
}
