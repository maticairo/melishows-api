package models

type Room struct {
	ID         string `json:"id"`
	RoomNumber int    `json:"room_number"`
	Seats      []Seat `json:"seats"`
}
