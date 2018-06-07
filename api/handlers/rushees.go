package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tomrom95/go-rush/api/models"
	data_models "github.com/tomrom95/go-rush/datastore/models"
	"github.com/tomrom95/go-rush/datastore/services"
)

func GetRushees(context *Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	rusheeService := services.NewRusheeService(context.DB)
	rushees, err := rusheeService.GetAllRushees()
	if err != nil {
		panic(err)
	}

	rusheeJSONList := make([]*models.Rushee, 0, len(rushees))
	for _, rushee := range rushees {
		rusheeJSON := &models.Rushee{}
		data_models.FillIn(rushee, rusheeJSON)
		rusheeJSONList = append(rusheeJSONList, rusheeJSON)
	}

	if err := json.NewEncoder(w).Encode(rusheeJSONList); err != nil {
		panic(err)
	}
}
