package models

type SeatPricing struct {
	ID    string `json:"id"`
	Price int    `json:"price"`
	Seats []Seat `json:"seats"`
}
