package models

import (
	"sort"
	"time"
)

type Show struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	Functions AllFunctions `json:"functions"`
}

type AllShows []Show

func (as AllShows) FindShow(showID string) *Show {
	for _, s := range as {
		if s.ID == showID {
			return &s
		}
	}
	return nil
}

func (as AllShows) FindByDate(dateFrom, dateTo time.Time) AllShows {
	var shows AllShows
	for _, s := range as {
		for _, f := range s.Functions {
			if (f.Date.After(dateFrom) || f.Date.Equal(dateFrom)) && (f.Date.Before(dateTo) || f.Date.Equal(dateTo)) {
				shows = append(shows, s)
			}
		}
	}
	return shows
}

func (as AllShows) FindByPrice(priceFrom, priceTo int) AllShows {
	if priceFrom == 0 && priceTo == 0 {
		return as
	}

	var shows AllShows
	for _, s := range as {
		for _, f := range s.Functions {
			for _, sp := range f.Pricing {
				if sp.Price >= priceFrom && sp.Price <= priceTo {
					shows = append(shows, s)
				}
			}
		}
	}
	return shows
}

func (as AllShows) OrderBy(orderKind string) AllShows {
	sort.Slice(as, func(i, j int) bool {
		if orderKind == "DESC" {
			return as[i].Name < as[j].Name
		} else {
			return as[i].Name > as[j].Name
		}
	})
	return as
}
