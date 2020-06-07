package main

import (
	"log"
	"net/http"

	"go-job-listing-aggregator/src/config/database"
	"go-job-listing-aggregator/src/listing"
	"go-job-listing-aggregator/src/query"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting go-job-listing-aggregator")

	router := mux.NewRouter()

	databaseConfig := database.NewDatabase()

	listing.NewDefaultModule(databaseConfig.DB, router)

	query.NewDefaultModule(databaseConfig.DB, router)

	log.Println("Started go-job-listing-aggregator")
	http.ListenAndServe(":9192", router)
}
