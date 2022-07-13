package helpers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/maticairo/melishows-api/pkg/models"
	"io/ioutil"
	"os"
)

func LoadData() (models.AllShows, []models.Theater) {
	//TODO load from file
	showsJsonFile, err := os.Open("db/shows.json")
	theatersJsonFile, err := os.Open("db/theaters.json")
	if err != nil {
		panic(err)
	}

	defer showsJsonFile.Close()
	defer theatersJsonFile.Close()

	showsByteValue, _ := ioutil.ReadAll(showsJsonFile)
	theatersByteValue, _ := ioutil.ReadAll(theatersJsonFile)

	var allShows models.AllShows
	var allTheaters []models.Theater

	err = json.Unmarshal(showsByteValue, &allShows)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(theatersByteValue, &allTheaters)

	if err != nil {
		panic(err)
	}

	return allShows, allTheaters

	/*
		theater := models.Theater{
			ID:    uuid.NewString(),
			Name:  "Teatro Colon",
			Rooms: generateTheaterRooms(),
		}

		shows := models.AllShows{
			{
				ID:   uuid.NewString(),
				Name: "El Lago de los Cisnes",
				Functions: []models.Function{
					{
						ID:            uuid.NewString(),
						Day:           "monday",
						StartingTime:  "",
						Duration:      120,
						Theater:       theater,
						TheaterRoomID: theater.Rooms[1].ID,
						Pricing:       nil,
					},
					{
						ID:            uuid.NewString(),
						Day:           "tuesday",
						StartingTime:  "",
						Duration:      120,
						Theater:       theater,
						TheaterRoomID: theater.Rooms[1].ID,
						Pricing: []models.SeatPricing{
							{
								ID:    uuid.NewString(),
								Price: 100,
								Seats: generateRoomSeats(),
							},
						},
					},
				},
			},
		}

		return shows*/
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
