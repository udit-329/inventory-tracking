package routes

import (
	"inventory-tracking/backend/pkg/controllers"

	"github.com/gorilla/mux"
)

//RegisterRoutes registers all the api routes to listen to.
func RegisterRoutes(dbHandler *controllers.BaseHandler, router *mux.Router) {
	router.HandleFunc("/add", dbHandler.AddItem).Methods("POST")
	router.HandleFunc("/get/{id}", dbHandler.GetItemByID).Methods("GET")
	router.HandleFunc("/get", dbHandler.GetItems).Methods("GET")
	router.HandleFunc("/update/{id}", dbHandler.UpdateItem).Methods("PUT")
	router.HandleFunc("/update/{id}", dbHandler.HandlePreFlightCors).Methods("OPTIONS")
	router.HandleFunc("/delete/{id}", dbHandler.DeleteItem).Methods("DELETE")
	router.HandleFunc("/delete/{id}", dbHandler.HandlePreFlightCors).Methods("OPTIONS")
	router.HandleFunc("/export", dbHandler.ExportItems).Methods("GET")
}
