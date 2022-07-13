package models

import "time"

type Booking struct {
	Dni        int      `json:"dni"`
	Name       string   `json:"name"`
	ShowID     string   `json:"show_id"`
	FunctionID string   `json:"function_id"`
	Seats      []string `json:"seats"`
}

type BookingInformation struct {
	Dni         int       `json:"dni"`
	Name        string    `json:"name"`
	ShowName    string    `json:"show_name"`
	TheaterName string    `json:"theater_name"`
	TheaterRoom int       `json:"theater_room"`
	Day         string    `json:"day"`
	ShowDate    time.Time `json:"show_date"`
	Seats       []string  `json:"seats"`
	TotalPrice  int       `json:"total_price"`
}
