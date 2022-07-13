package models

import "time"

type SearchCriteria struct {
	DateFrom  time.Time `json:"date_from"`
	DateTo    time.Time `json:"date_to"`
	PriceFrom int       `json:"price_from"`
	PriceTo   int       `json:"price_to"`
	Order     string    `json:"order"`
}
