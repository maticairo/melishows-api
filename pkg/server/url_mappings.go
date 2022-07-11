package server

import (
	"github.com/maticairo/melishows-api/pkg/melishows"
	"net/http"
)

func MapURLs() {

	http.Handle("/", new(melishows.MelishowsHandler))
}
