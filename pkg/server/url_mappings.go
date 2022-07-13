package server

import (
	"github.com/gorilla/mux"
	"github.com/maticairo/melishows-api/pkg/melishows"
	"github.com/maticairo/melishows-api/pkg/middlewares"
	"log"
	"net/http"
)

func MapURLs() *mux.Router {

	controller := melishows.NewController()

	router := mux.NewRouter()

	/** GCP utils */
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		middlewares.ResponseWithJSON(w, "Hi Mate! Perhaps you may want to read our Readme.md file ;)")
	})
	router.HandleFunc("/_ah/warmup", func(w http.ResponseWriter, r *http.Request) {
		log.Println("warmup done")
	})
	router.HandleFunc("/ping", middlewares.Ping)

	/* API endpoint */
	router.HandleFunc("/shows/all", controller.GetAllShows)
	router.HandleFunc("/shows/search",
		controller.SearchShows).Queries(
		"date_from", "{dateFrom}",
		"date_to", "{dateTo}",
		"price_from", "{priceFrom}",
		"price_to", "{priceTo}",
		"order_kind", "{orderKind}")

	router.HandleFunc("/availableSeats",
		controller.GetAvailableSeats).Queries(
		"show_id", "{showID}",
		"function_id", "{functionID}")

	router.HandleFunc("/book",
		controller.BookSeats).Methods("POST")

	return router
}
