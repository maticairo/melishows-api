package melishows

import (
	"github.com/maticairo/melishows-api/pkg/helpers"
	"github.com/maticairo/melishows-api/pkg/models"
)

type Repository struct {
	allShows     models.AllShows
	allTheaters  []models.Theater
	reservations []models.BookingInformation
}

func NewRepository() *Repository {
	shows, theaters := helpers.LoadData()
	repository := Repository{
		allShows:    shows,
		allTheaters: theaters,
	}
	return &repository
}

func (r Repository) GetAllShows() *models.AllShows {
	return &r.allShows
}

func (r Repository) GetAllTheaters() *[]models.Theater {
	return &r.allTheaters
}

func (r Repository) UpdateAllShows(allShows models.AllShows) {
	r.allShows = allShows
}

func (r Repository) SaveReservation(information models.BookingInformation) {
	r.reservations = append(r.reservations, information)
}
