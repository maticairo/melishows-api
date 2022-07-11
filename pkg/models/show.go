package models

type Show struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Functions []Function `json:"functions"`
}
