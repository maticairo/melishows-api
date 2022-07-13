package middlewares

import (
	"encoding/json"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("pong")
	if err != nil {
		panic(err)
	}
}

func ResponseWithJSON(w http.ResponseWriter, v any, responseStatus ...int) {
	w.Header().Set("Content-Type", "application/json")

	if len(responseStatus) > 0 {
		w.WriteHeader(responseStatus[0])
	} else {
		w.WriteHeader(http.StatusOK)
	}

	err := json.NewEncoder(w).Encode(v)

	if err != nil {
		return
	}

}
