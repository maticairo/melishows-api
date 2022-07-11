package models

type Response struct {
	Theater Theater `json:"theater"`
	Show    Show    `json:"show"`
}
