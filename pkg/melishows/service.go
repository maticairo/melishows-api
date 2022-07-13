package melishows

import (
	"errors"
	"fmt"
	"github.com/maticairo/melishows-api/pkg/cache"
	"github.com/maticairo/melishows-api/pkg/models"
	"time"
)

type Service struct {
	repository *Repository
	cache      *cache.Cache
}

var allShowsCacheKey = "allShows"

func NewService() *Service {
	return &Service{
		cache:      cache.NewCache(),
		repository: NewRepository(),
	}
}

func (s Service) GetAllShows() *models.AllShows {
	allShowsFromCache := s.cache.Get(allShowsCacheKey)

	if allShowsFromCache == nil {
		allShows := s.repository.GetAllShows()
		s.cache.Set(allShowsCacheKey, allShows, 1*time.Minute)
		return allShows
	}

	return allShowsFromCache.(*models.AllShows)
}

func (s Service) GetAvailableSeats(showID, functionID string) models.AvailableSeats {
	allShows := s.repository.GetAllShows()

	response := models.AvailableSeats{
		Show:     showID,
		Function: functionID,
		Seats:    nil,
	}

	show := allShows.FindShow(showID)
	function := show.Functions.FindFunction(functionID)

	for _, sp := range function.Pricing {
		var seats []models.Seat
		var seatPricing models.SeatPricing
		seatPricing.ID = sp.ID
		seatPricing.Price = sp.Price
		for _, s := range sp.Seats {
			if !s.Booked {
				seats = append(seats, s)
			}
		}
		seatPricing.Seats = seats
		response.Seats = append(response.Seats, seatPricing)
	}

	return response
}

func (s Service) BookSeats(booking models.Booking) (*models.BookingInformation, error) {
	allShows := s.repository.GetAllShows()
	fmt.Println(booking)
	response := &models.BookingInformation{
		Dni:        booking.Dni,
		Name:       booking.Name,
		TotalPrice: 0,
	}

	for _, show := range *allShows {
		if show.ID == booking.ShowID {
			response.ShowName = show.Name
			for _, function := range show.Functions {
				if function.ID == booking.FunctionID {
					theater := s.getTheaterInformation(function.TheaterID)
					response.TheaterName = theater.Name
					response.TheaterRoom = s.getTheaterRoom(function.TheaterRoomID, theater)
					response.Day = function.Day
					response.StartingTime = function.StartingTime
					for _, seatToBook := range booking.Seats {
						for _, sp := range function.Pricing {
							var bookedSeats = 0
							var seatsToBook []models.Seat
							for _, s := range sp.Seats {
								if !s.Booked {
									if fmt.Sprintf("%d-%s", s.RowNumber, s.Identifier) == seatToBook {
										seatsToBook = append(seatsToBook, s)
										s.Booked = true
										bookedSeats++
									}
								} else {
									return nil, errors.New("some Seats are already booked, please try again with other seats")
								}
							}
							response.TotalPrice += bookedSeats * sp.Price
							response.Seats = append(response.Seats, seatToBook)
						}
					}
				}
			}
		}
	}
	allShows = s.updateSeats(booking.ShowID, booking.FunctionID, allShows, response.Seats)
	s.cache.Update(allShowsCacheKey, allShows, 1*time.Minute)
	s.repository.UpdateAllShows(allShows)
	s.repository.SaveReservation(*response)
	return response, nil
}

func (s Service) updateSeats(showID string, functionID string, allShows *models.AllShows, seats []string) *models.AllShows {
	show := allShows.FindShow(showID)
	function := show.Functions.FindFunction(functionID)
	for _, seat := range seats {
		for _, sp := range function.Pricing {
			for _, s := range sp.Seats {
				if fmt.Sprintf("%d-%s", s.RowNumber, s.Identifier) == seat {
					s.Booked = true
				}
			}
		}
	}
	return allShows
}

func (s Service) getTheaterInformation(theaterID string) models.Theater {
	theaters := s.repository.GetAllTheaters()
	for _, theater := range *theaters {
		if theater.ID == theaterID {
			return theater
		}
	}
	return models.Theater{}
}

func (s Service) getTheaterRoom(theaterRoomID string, theater models.Theater) int {
	for _, room := range theater.Rooms {
		if room.ID == theaterRoomID {
			return room.RoomNumber
		}
	}
	return 0
}

func (s Service) Search(searchCriteria models.SearchCriteria) models.AllShows {
	allShows := s.repository.GetAllShows()
	shows := allShows.FindByDate(searchCriteria.DateFrom, searchCriteria.DateTo)
	shows = shows.FindByPrice(searchCriteria.PriceFrom, searchCriteria.PriceTo)
	shows.OrderBy(searchCriteria.Order)
	return shows
}
