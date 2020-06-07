package github

import (
	"go-job-listing-aggregator/src/listing"
	"go-job-listing-aggregator/src/query"
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
