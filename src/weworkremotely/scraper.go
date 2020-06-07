package weworkremotely

import (
	"go-job-listing-aggregator/src/listing"
	"go-job-listing-aggregator/src/query"
	"strings"
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

func containsAnyString(text string, tokens []string) bool {
	for _, token := range tokens {
		if strings.Contains(text, token) {
			return true
		}
	}

	return false
}

func isPositionListingRelevant(
	keywords []string,
	positionListing *jobListingRSSModel) bool {
	var keywordsLowerCase []string
	for _, word := range keywords {
		keywordsLowerCase = append(keywordsLowerCase, strings.ToLower(word))
	}

	lowerCaseText := strings.ToLower(positionListing.Title)
	if containsAnyString(lowerCaseText, keywordsLowerCase) {
		return true
	}

	lowerCaseText = strings.ToLower(positionListing.Description)
	if containsAnyString(lowerCaseText, keywordsLowerCase) {
		return true
	}

	return false
}

func filterNonRelevantPositionListings(
	keywords []string,
	positionListings []jobListingRSSModel) []jobListingRSSModel {
	var result []jobListingRSSModel

	for _, listing := range positionListings {
		if isPositionListingRelevant(keywords, &listing) {
			result = append(result, listing)
		}
	}

	return result
}
