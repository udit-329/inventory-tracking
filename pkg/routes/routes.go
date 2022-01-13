package routes

import (
	"inventory-tracking/pkg/controllers"

	"github.com/gorilla/mux"
)

//RegisterRoutes registers all the api routes to listen to.
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/add", controllers.AddItem).Methods("POST")
	router.HandleFunc("/get/{id}", controllers.GetItemByID).Methods("GET")
	router.HandleFunc("/get", controllers.GetItems).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdateItem).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeleteItem).Methods("DELETE")
	router.HandleFunc("/export", controllers.ExportItems).Methods("GET")
}
