package jobsbg

import (
	"go-job-listing-aggregator/src/listing"
	"go-job-listing-aggregator/src/query"
)

const domain = "https://www.jobs.bg/"

// Scraper accesses the jobs.bg website and searches for
// job listings. Found entries are transformed to listing.Listing
// structs and persisted.
type Scraper struct {
	listingService listing.Service
	queryService   query.Service
}

// NewScraper returns a new instance.
func NewScraper(listingService listing.Service, queryService query.Service) Scraper {
	return Scraper{listingService: listingService, queryService: queryService}
}
