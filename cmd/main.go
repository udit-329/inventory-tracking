package main

import (
	"inventory-tracking/backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	mux.CORSMethodMiddleware(router)
	routes.RegisterRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
