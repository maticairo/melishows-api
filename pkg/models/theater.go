package models

type Theater struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Rooms []Room `json:"rooms"`
}

func (t Theater) Room(roomNumber int) *Room {
	for _, r := range t.Rooms {
		if r.RoomNumber == roomNumber {
			return &r
		}
	}
	return nil
}
