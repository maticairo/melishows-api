package tests

import (
	"fmt"
	"github.com/maticairo/melishows-api/pkg/melishows"
	"github.com/maticairo/melishows-api/pkg/models"
	"reflect"
	"testing"
	"time"
)

func TestIntegration_Book(t *testing.T) {
	functionDate, _ := time.Parse(time.RFC3339, time.Date(2020, 01, 01, 15, 00, 00, 00, time.UTC).Format(time.RFC3339))
	cases := []struct {
		name             string
		service          *melishows.Service
		bookingRequest   models.Booking
		expectedResponse *models.BookingInformation
		expectedErr      bool
	}{
		{
			name:    "booking_integration_test/OK",
			service: melishows.NewService(),
			bookingRequest: models.Booking{
				Dni:        1,
				Name:       "name",
				ShowID:     "62963928-f501-4af3-bafd-0acce2321668",
				FunctionID: "7c336060-02c7-4d30-aec0-507e7b4c4c40",
				Seats: []string{
					"1-A",
					"1-B",
					"1-C",
				},
			},
			expectedResponse: &models.BookingInformation{
				Dni:         1,
				Name:        "name",
				ShowName:    "El Lago de los Cisnes",
				TheaterName: "Teatro Colon",
				TheaterRoom: 2,
				Date:        functionDate,
				Seats: []string{
					"1-A",
					"1-B",
					"1-C",
				},
				TotalPrice: 300,
			},
		},
		{
			name:    "booking_integration_test/SeatAlreadyBooked",
			service: melishows.NewService(),
			bookingRequest: models.Booking{
				Dni:        1,
				Name:       "name",
				ShowID:     "62963928-f501-4af3-bafd-0acce2321668",
				FunctionID: "7c336060-02c7-4d30-aec0-507e7b4c4c40",
				Seats: []string{
					"1-A",
				},
			},
			expectedResponse: nil,
			expectedErr:      true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			response, err := c.service.BookSeats(c.bookingRequest)

			if err != nil && !c.expectedErr {
				t.Error(fmt.Sprintf("Unexpected error %v when unexpected", err.Error()))
			}

			if err == nil && c.expectedErr {
				t.Error(fmt.Sprintf("No error when expected"))
			}

			if !reflect.DeepEqual(response, c.expectedResponse) {
				t.Error(fmt.Sprintf("Mismatch responses. Actual: %v - Expected: %v", response, c.expectedResponse))
			}
		})
	}

}
