package stackoverflow

import (
	"go-job-listing-aggregator/src/listing"
	"go-job-listing-aggregator/src/query"
)

// Scraper extracts listings from the StackOverflow rss feed.
type Scraper struct {
	listingService listing.Service
	queryService   query.Service
}

// NewScraper returns a new instance.
func NewScraper(listingService listing.Service, queryService query.Service) Scraper {
	return Scraper{listingService: listingService, queryService: queryService}
}
