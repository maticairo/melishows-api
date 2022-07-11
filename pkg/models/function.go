package models

type Function struct {
	ID           string        `json:"id"`
	Day          string        `json:"day"`
	StartingTime string        `json:"starting_time"`
	Duration     int           `json:"duration"`
	Theater      Theater       `json:"theater"`
	TheaterRoom  Room          `json:"theater_room"`
	Pricing      []SeatPricing `json:"pricing"`
}
