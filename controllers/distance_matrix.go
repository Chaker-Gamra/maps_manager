package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Chaker-Gamra/maps_manager/models"
)

//  Get Distance matrix godoc
// @Summary Calculate Matrix
// @Description Get a 2D array of directions
// @Tags Distance Matrix API
// @Accept  json
// @Produce  json
// @Success 200
// @Param provider query string false "Select A Provider" Enums(Google, Here)
// @Param origins query string false "Put athe list of Origins in this format lat,lng|lat,lng ..."
// @Param destinations query string false "Put the list of Destinations in this format lat,lng|lat,lng ..."
// @Param traffic query string false "Traffic Option" Enums(True, False)
// @Success 200 {object} models.DistanceMatrix
// @Router /distance_matrix [get]
func CalculateMatrix(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("distance matrix API")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.DistanceMatrix{})
}
