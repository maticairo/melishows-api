package melishows

import (
	"github.com/maticairo/melishows-api/pkg/models"
	"testing"
)

func TestNewRepository(t *testing.T) {
	t.Run("TestNewRepository", func(t *testing.T) {
		NewRepository()
	})
}

func TestRepository_GetAllShows(t *testing.T) {
	expectedShows := models.AllShows{
		{
			ID: "62963928-f501-4af3-bafd-0acce2321668",
		},
	}
	repository := NewRepository()
	t.Run("RepositoryGetAllShows", func(t *testing.T) {
		shows := repository.GetAllShows()
		if (*shows)[0].ID != expectedShows[0].ID {
			t.Error("Unexpected shows from repository")
		}
	})
}

func TestRepository_GetAllTheaters(t *testing.T) {
	expectedTheaters := []models.Theater{
		{
			ID: "1eed8488-bac2-466c-95a8-cc6c450082b5",
		},
	}
	repository := NewRepository()
	t.Run("RepositoryGetAllTheaters", func(t *testing.T) {
		theaters := repository.GetAllTheaters()
		if (*theaters)[0].ID != expectedTheaters[0].ID {
			t.Error("Unexpected theaters from repository")
		}
	})
}

func TestRepository_SaveReservation(t *testing.T) {
	reservation := models.BookingInformation{
		Dni: 1,
	}
	repository := NewRepository()
	t.Run("RepositorySaveReservation", func(t *testing.T) {
		repository.SaveReservation(reservation)
		if len(repository.reservations) < 1 {
			t.Error("No reservation saved")
		}
	})
}

func TestRepository_UpdateAllShows(t *testing.T) {
	repository := NewRepository()
	t.Run("RepositoryUpdateAllShows", func(t *testing.T) {
		shows := repository.GetAllShows()
		for _, show := range *shows {
			show.Name = "UpdatedName"
		}
		repository.UpdateAllShows(*shows)

		repository.GetAllShows()
		for _, show := range *shows {
			if show.Name != "UpdatedName" {
				t.Error("Error updating shows")
			}
		}
	})
}
