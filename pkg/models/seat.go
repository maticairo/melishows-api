package models

type Seat struct {
	RowNumber  int    `json:"row_number"`
	Identifier string `json:"identifier"`
}
