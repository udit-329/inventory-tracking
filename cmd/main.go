package main

import (
	"inventory-tracking/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
