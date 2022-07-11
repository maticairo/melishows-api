package main

import (
	"github.com/maticairo/melishows-api/pkg/server"
	"log"
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.MapURLs()
	log.Fatal(s.ListenAndServe())
}
