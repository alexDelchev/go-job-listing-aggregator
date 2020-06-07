package github

import (
	"go-job-listing-aggregator/src/listing"
	"go-job-listing-aggregator/src/query"
	"log"
	"sync"
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

// RunForActiveQueries calls Scrape in paralel for all active queries
func (s *Scraper) RunForActiveQueries() {
	queries, err := s.queryService.GetActiveQueries()
	if err != nil {
		log.Println("Error in while fetching active queries")
		return
	}

	var waitGroup sync.WaitGroup

	for _, searchQuery := range queries {
		waitGroup.Add(1)

		go func(searchQuery query.Query) {
			defer waitGroup.Done()
			s.Scrape(searchQuery)
		}(searchQuery)
	}

	waitGroup.Wait()
}
