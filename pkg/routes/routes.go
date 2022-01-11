package routes

import (
	"inventory-tracking/inventory-tracking/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/add/", controllers.AddItem).Methods("POST")
	router.HandleFunc("/get/{id}", controllers.GetItemById).Methods("GET")
	router.HandleFunc("/get/", controllers.GetItems).Methods("GET")
	router.HandleFunc("/update/{SKU}", controllers.UpdateItem).Methods("PUT")
	router.HandleFunc("/delete/{SKU}", controllers.DeleteItem).Methods("DELETE")
	router.HandleFunc("/export/", controllers.ExportItems).Methods("GET")
}
