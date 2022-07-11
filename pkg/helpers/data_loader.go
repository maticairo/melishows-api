package helpers

import (
	"github.com/google/uuid"
	"github.com/maticairo/melishows-api/pkg/models"
)

func LoadData() (models.Theater, models.Show) {
	theater := models.Theater{
		ID:    uuid.NewString(),
		Name:  "Teatro Colon",
		Rooms: generateTheaterRooms(),
	}

	show := models.Show{
		ID:   uuid.NewString(),
		Name: "El Lago de los Cisnes",
		Functions: []models.Function{
			{
				ID:           uuid.NewString(),
				Day:          "monday",
				StartingTime: "",
				Duration:     120,
				Theater:      theater,
				TheaterRoom:  *theater.Room(1),
				Pricing:      nil,
			},
		},
	}

	return theater, show
}

func generateTheaterRooms() []models.Room {
	rooms := make([]models.Room, 0)
	for i := 1; i < 5; i++ {
		r := models.Room{
			ID:         uuid.NewString(),
			RoomNumber: i,
			Seats:      generateRoomSeats(),
		}
		rooms = append(rooms, r)
	}
	return rooms
}

func generateRoomSeats() []models.Seat {
	seatIdentifiers := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I",
	}
	seats := make([]models.Seat, 0)
	for i := 1; i < 15; i++ {
		for j := 0; j <= 8; j++ {
			s := models.Seat{
				RowNumber:  i,
				Identifier: seatIdentifiers[j],
			}
			seats = append(seats, s)
		}
	}
	return seats
}
