package melishows

import (
	"encoding/json"
	"github.com/maticairo/melishows-api/pkg/middlewares"
	"github.com/maticairo/melishows-api/pkg/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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

/*
	GetAllShows returns a list with all the information about the shows
*/
func (m *Controller) GetAllShows(w http.ResponseWriter, r *http.Request) {
	allShows := m.service.GetAllShows()
	middlewares.ResponseWithJSON(w, allShows)
	return
}

/*
	GetAvailableSeats returns a list of the available seats given a showID and a functionID
*/
func (m *Controller) GetAvailableSeats(w http.ResponseWriter, r *http.Request) {
	showID := r.URL.Query().Get("show_id")
	functionID := r.URL.Query().Get("function_id")

	if showID == "" || functionID == "" {
		middlewares.ResponseWithJSON(w, "You must provide both show_id and function_id query params", http.StatusBadRequest)
		return
	}

	availableSeats, err := m.service.GetAvailableSeats(showID, functionID)

	if err != nil {
		middlewares.ResponseWithJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	middlewares.ResponseWithJSON(w, availableSeats)
	return
}

/*
	BookSeats makes a reservation of a given set of seats and returns a printable ticket
*/
func (m *Controller) BookSeats(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		middlewares.ResponseWithJSON(w, err.Error(), http.StatusCreated)
		return
	}

	err = json.Unmarshal(reqBody, &booking)

	if err != nil {
		middlewares.ResponseWithJSON(w, err.Error(), http.StatusCreated)
		return
	}

	bookingInformation, err := m.service.BookSeats(booking)

	if err != nil {
		middlewares.ResponseWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(bookingInformation.Seats) == 0 {
		middlewares.ResponseWithJSON(w, "No seats booked, please check you provided valid seats", http.StatusNoContent)
		return
	}

	middlewares.ResponseWithJSON(w, *bookingInformation, http.StatusCreated)
	return
}

/*
	SearchShows makes a deep search given the following params:
		- Date from
		- Date To
		- Price from
		- Price To
		- Order Kind
*/
func (m *Controller) SearchShows(w http.ResponseWriter, r *http.Request) {
	strDateFrom := r.URL.Query().Get("date_from")
	strDateTo := r.URL.Query().Get("date_to")

	if strDateFrom == "" || strDateTo == "" {
		middlewares.ResponseWithJSON(w, "date_from and date_to must not be empty", http.StatusBadRequest)
		return
	}

	dateFrom, err := time.Parse(time.RFC3339, strDateFrom)
	if err != nil {
		middlewares.ResponseWithJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	dateTo, err := time.Parse(time.RFC3339, strDateTo)

	if err != nil {
		middlewares.ResponseWithJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	priceFrom, err := strconv.Atoi(r.URL.Query().Get("price_from"))

	if err != nil {
		middlewares.ResponseWithJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	priceTo, err := strconv.Atoi(r.URL.Query().Get("price_to"))

	if err != nil {
		middlewares.ResponseWithJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if priceTo == 0 {
		priceTo = priceFrom
	}

	if priceTo < 0 || priceTo < priceFrom {
		middlewares.ResponseWithJSON(w, "price input is invalid", http.StatusBadRequest)
		return
	}

	orderKind := r.URL.Query().Get("order_kind")

	if orderKind == "" {
		orderKind = "ASC"
	}

	if orderKind != "ASC" && orderKind != "DESC" {
		middlewares.ResponseWithJSON(w, "order kind must be 'asc' or 'desc'", http.StatusBadRequest)
		return
	}

	searchCriteria := models.SearchCriteria{
		DateFrom:  dateFrom,
		DateTo:    dateTo,
		PriceFrom: priceFrom,
		PriceTo:   priceTo,
		OrderKind: orderKind,
	}

	searchInformation := m.service.SearchShows(searchCriteria)

	middlewares.ResponseWithJSON(w, searchInformation)
	return
}

func (m *Controller) Reset(w http.ResponseWriter, r *http.Request) {

	m.service.Reset()

	middlewares.ResponseWithJSON(w, "{status: ok}")
	return
}
