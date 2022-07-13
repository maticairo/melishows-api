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

	cases := []struct {
		name             string
		service          *melishows.Service
		bookingRequest   models.Booking
		expectedResponse models.BookingInformation
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
			expectedResponse: models.BookingInformation{
				Dni:         0,
				Name:        "",
				ShowName:    "",
				TheaterName: "",
				TheaterRoom: 0,
				Date:        time.Time{},
				Seats:       nil,
				TotalPrice:  0,
			},
		},
		{
			name:    "booking_integration_test/FAIL - cached",
			service: melishows.NewService(),
			bookingRequest: models.Booking{
				Dni:        0,
				Name:       "",
				ShowID:     "",
				FunctionID: "",
				Seats:      nil,
			},
			expectedResponse: models.BookingInformation{
				Dni:         0,
				Name:        "",
				ShowName:    "",
				TheaterName: "",
				TheaterRoom: 0,
				Date:        time.Time{},
				Seats:       nil,
				TotalPrice:  0,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			response, err := c.service.BookSeats(c.bookingRequest)

			if err != nil && !c.expectedErr {
				t.Error(fmt.Sprintf("Unexpected error %v when unexpected", err.Error()))
			}

			if err == nil && c.expectedErr {
				t.Error(fmt.Sprintf("No error %v when expected", err.Error()))
			}

			if !reflect.DeepEqual(response, c.expectedResponse) {
				t.Error(fmt.Sprintf("Mismatch responses. Actual: %v - Expected: %v", response, c.expectedResponse))
			}
		})
	}

}
