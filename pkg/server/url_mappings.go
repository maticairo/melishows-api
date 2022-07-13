package server

import (
	"github.com/gorilla/mux"
	"github.com/maticairo/melishows-api/pkg/melishows"
	"github.com/maticairo/melishows-api/pkg/middlewares"
)

func MapURLs() *mux.Router {

	controller := melishows.NewController()

	router := mux.NewRouter()
	router.HandleFunc("/ping", middlewares.Ping)
	router.HandleFunc("/getAllShows", controller.GetAllShows)
	//router.HandleFunc("/searchShows", melishows.ReadSomething).Methods("POST")
	router.HandleFunc("/getAvailableSeats",
		controller.GetAvailableSeats).Queries(
		"show_id", "{showID}",
		"function_id", "{function_id}")
	router.HandleFunc("/book",
		controller.BookSeats).Methods("POST")

	return router
}
