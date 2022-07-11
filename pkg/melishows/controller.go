package melishows

import (
	"encoding/json"
	"github.com/maticairo/melishows-api/pkg/helpers"
	"github.com/maticairo/melishows-api/pkg/models"
	"net/http"
)

type MelishowsHandler struct {
}

func (m MelishowsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	theater, show := helpers.LoadData()
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(models.Response{
		Theater: theater,
		Show:    show,
	})

	if err != nil {
		return
	}
}
