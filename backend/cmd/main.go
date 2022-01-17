package main

import (
	"inventory-tracking/backend/pkg/config"
	"inventory-tracking/backend/pkg/controllers"
	"inventory-tracking/backend/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	mux.CORSMethodMiddleware(router)

	db := config.Connect()
	dbHandler := controllers.NewBaseHandler(db)

	routes.RegisterRoutes(dbHandler, router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
