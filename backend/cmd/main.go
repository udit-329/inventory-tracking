package main

import (
	"inventory-tracking/backend/pkg/config"
	"inventory-tracking/backend/pkg/controllers"
	"inventory-tracking/backend/pkg/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	mux.CORSMethodMiddleware(router)

	db := config.Connect()
	dbHandler := controllers.NewBaseHandler(db)

	port := os.Getenv("PORT")
	routes.RegisterRoutes(dbHandler, router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
