package router

import (
	"net/http"

	"github.com/Chaker-Gamra/maps_manager/controllers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {

	//Create a new mux router
	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/directions", controllers.CalculateRoute).Methods("GET")       //GET / directions
	r.HandleFunc("/distance_matrix", controllers.CalculateMatrix).Methods("GET") //GET / distance_matrix

	//Serve static files
	sh := http.StripPrefix("/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/").Handler(sh)

	return r
}
