package models

type AvailableSeats struct {
	Show     string        `json:"show"`
	Function string        `json:"function"`
	Seats    []SeatPricing `json:"seats"`
}
