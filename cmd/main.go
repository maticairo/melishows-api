package main

import (
	"fmt"
	"github.com/maticairo/melishows-api/pkg/server"
	"log"
	"net/http"
	"os"
)

func main() {
	router := server.MapURLs()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
