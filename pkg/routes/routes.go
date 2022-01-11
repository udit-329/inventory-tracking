package routes

import (
	"inventory-tracking/inventory-tracking/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/add/", controllers.AddItem).Methods("POST")
	router.HandleFunc("/update/{SKU}", controllers.AddItem).Methods("POST") //PUT
	router.HandleFunc("/delete/{SKU}", controllers.AddItem).Methods("DELETE")
}
