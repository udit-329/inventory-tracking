package routes

import (
	"inventory-tracking/backend/pkg/controllers"

	"github.com/gorilla/mux"
)

//RegisterRoutes registers all the api routes to listen to.
func RegisterRoutes(dbHandler *controllers.BaseHandler, router *mux.Router) {
	router.HandleFunc("/add", dbHandler.AddItem).Methods("POST")
	router.HandleFunc("/get", dbHandler.GetItem).Methods("GET")
	router.HandleFunc("/update", dbHandler.UpdateItem).Methods("PUT")
	router.HandleFunc("/update", dbHandler.HandlePreFlightCors).Methods("OPTIONS")
	router.HandleFunc("/delete", dbHandler.DeleteItem).Methods("DELETE")
	router.HandleFunc("/delete", dbHandler.HandlePreFlightCors).Methods("OPTIONS")
	router.HandleFunc("/export", dbHandler.ExportItems).Methods("GET")
}
