package main

import (
	"log"
	"net/http"

	"go-job-listing-aggregator/src/config/database"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting go-job-listing-aggregator")

	router := mux.NewRouter()

	database.NewDatabase()

	log.Println("Started go-job-listing-aggregator")
	http.ListenAndServe(":9192", router)
}
