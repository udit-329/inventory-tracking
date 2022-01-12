package main

import (
	"inventory-tracking/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "fmt"
	// "os"
	// "github.com/joho/godotenv"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
