package melishows

import (
	"encoding/json"
	"github.com/maticairo/melishows-api/pkg/middlewares"
	"github.com/maticairo/melishows-api/pkg/models"
	"io/ioutil"
	"net/http"
)

type Controller struct {
	service *Service
}

func NewController() *Controller {
	controller := Controller{
		service: NewService(),
	}
	return &controller
}

func (m *Controller) GetAllShows(w http.ResponseWriter, r *http.Request) {
	allShows := m.service.GetAllShows()
	middlewares.ResponseWithJSON(w, allShows)
}

func ReadSomething(w http.ResponseWriter, r *http.Request) {
	var show models.Booking
	reqBody, err := ioutil.ReadAll(r.Body)

	if err == nil {
		return
	}

	err = json.Unmarshal(reqBody, &show)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(show)

	if err != nil {
		return
	}

	return
}

func (m *Controller) GetAvailableSeats(w http.ResponseWriter, r *http.Request) {
	showID := r.FormValue("showID")
	functionID := r.FormValue("functionID")

	availableSeats := m.service.GetAvailableSeats(showID, functionID)

	middlewares.ResponseWithJSON(w, availableSeats)

}

func (m *Controller) BookSeats(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(reqBody, &booking)

	if err != nil {
		panic(err)
	}

	bookingInformation, err := m.service.BookSeats(booking)

	if err != nil {
		middlewares.ResponseWithJSON(w, err, http.StatusInternalServerError)
	}

	middlewares.ResponseWithJSON(w, *bookingInformation, http.StatusCreated)

}
