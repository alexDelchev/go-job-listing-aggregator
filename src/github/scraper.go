package github

import (
	"go-job-listing-aggregator/src/listing"
	"go-job-listing-aggregator/src/query"
	"log"
)

// Scraper contains a listing service and query service
type Scraper struct {
	listingService listing.Service
	queryService   query.Service
}

// NewScraper returns a new instance
func NewScraper(listingService listing.Service, queryService query.Service) Scraper {
	return Scraper{listingService: listingService, queryService: queryService}
}

// Scrape fetches job listings from github for the given Query
func (s *Scraper) Scrape(searchQuery query.Query) {
	log.Printf("Starting for Query %+v", searchQuery)

	results, err := searchPositions(searchQuery.Keywords, searchQuery.Location)
	if err != nil {
		return
	}
	resultsTransformed := transformToListingModelSlice(searchQuery.ID, results)

	s.listingService.CreateListings(resultsTransformed)
}
