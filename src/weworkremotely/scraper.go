package weworkremotely

import (
	"go-job-listing-aggregator/src/listing"
	"go-job-listing-aggregator/src/query"
)

// Scraper loads the WeWorkRemotely rss feed and filters out
// job listings based on keywords. The resulting listings are
// persisted to the database.
type Scraper struct {
	listingService listing.Service
	queryService   query.Service
}

// NewScraper returns a new intance.
func NewScraper(listingService listing.Service, queryService query.Service) Scraper {
	return Scraper{listingService: listingService, queryService: queryService}
}
