package github

import (
	"go-job-listing-aggregator/src/listing"
	"go-job-listing-aggregator/src/query"

	"github.com/gorilla/mux"
)

// Module contains a scraper, scheduler and controller,
// with the Scraper and Scheduler being exported.
type Module struct {
	Scraper    Scraper
	Scheduler  Scheduler
	controller controller
}

// NewDefaultModule injects the default dependencies.
func NewDefaultModule(
	listingService listing.Service,
	queryService query.Service,
	router *mux.Router) Module {

	scraper := NewScraper(listingService, queryService)
	scheduler := NewScheduler(&scraper)
	controller := newContoller(&scheduler, router)

	return Module{Scraper: scraper, Scheduler: scheduler, controller: controller}
}
