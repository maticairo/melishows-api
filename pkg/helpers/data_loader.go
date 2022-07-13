package helpers

import (
	"encoding/json"
	"github.com/maticairo/melishows-api/pkg/models"
	"io/ioutil"
	"os"
)

func LoadData() (models.AllShows, []models.Theater) {
	showsJsonFile, err := os.Open("db/shows.json")

	if err != nil {
		panic(err)
	}

	theatersJsonFile, err := os.Open("db/theaters.json")
	if err != nil {
		panic(err)
	}

	defer showsJsonFile.Close()
	defer theatersJsonFile.Close()

	showsByteValue, _ := ioutil.ReadAll(showsJsonFile)
	theatersByteValue, _ := ioutil.ReadAll(theatersJsonFile)

	var allShows models.AllShows
	var allTheaters []models.Theater

	err = json.Unmarshal(showsByteValue, &allShows)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(theatersByteValue, &allTheaters)

	if err != nil {
		panic(err)
	}

	return allShows, allTheaters
}
