package stackoverflow

import (
	"go-job-listing-aggregator/src/listing"
	"go-job-listing-aggregator/src/query"
	"log"
	"sync"
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

// Scrape extracts listings from the rss feed for the given serach Query.
func (s *Scraper) Scrape(searchQuery query.Query) {
	log.Printf("Starting for Query %+v", searchQuery)

	results, err := searchPositions(searchQuery.Keywords, searchQuery.Location)
	if err != nil {
		return
	}
	resultsTransformed := transformToListingModelSlice(searchQuery.ID, results)

	s.listingService.CreateListings(resultsTransformed)
}

// RunForActiveQueries calls Scraper concurrently for all active queries.
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
