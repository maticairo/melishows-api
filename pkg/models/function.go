package models

type Function struct {
	ID            string        `json:"id"`
	Day           string        `json:"day"`
	StartingTime  string        `json:"starting_time"`
	Duration      int           `json:"duration"`
	TheaterID     string        `json:"theater_id"`
	TheaterRoomID string        `json:"theater_room_id"`
	Pricing       []SeatPricing `json:"pricing"`
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
