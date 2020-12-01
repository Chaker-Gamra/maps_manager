package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Chaker-Gamra/maps_manager/models"
	"github.com/gorilla/mux"
)

// GetDirections godoc
// @Summary Calculate Route
// @Description Our First API To Get Direction info between two locations
// @Tags direction API
// @Accept  json
// @Produce  json
// @Param provider query string false "Select A Provider" Enums(Google, Here)
// @Param origin query string false "Put An Origin in this format lat,lng"
// @Param destination query string false "Put A Destination in this format lat,lng"
// @Param withPath query string false "WithPath Option" Enums(True, False)
// @Param traffic query string false "Traffic Option" Enums(True, False)
// @Success 200 {object} models.Direction
// @Router /directions [get]
func CalculateRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("directions API")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Print(params)
	json.NewEncoder(w).Encode(models.Direction{})
}
