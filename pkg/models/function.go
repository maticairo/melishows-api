package models

import "time"

type Function struct {
	ID            string         `json:"id"`
	Date          time.Time      `json:"date"`
	Duration      int            `json:"duration"`
	TheaterID     string         `json:"theater_id"`
	TheaterRoomID string         `json:"theater_room_id"`
	Pricing       []*SeatPricing `json:"pricing"`
}

type AllFunctions []Function

func (af AllFunctions) FindFunction(functionID string) *Function {
	for _, f := range af {
		if f.ID == functionID {
			return &f
		}
	}
	return nil
}
