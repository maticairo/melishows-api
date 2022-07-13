package models

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
