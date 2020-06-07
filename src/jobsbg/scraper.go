package jobsbg

import (
	"go-job-listing-aggregator/src/listing"
	"go-job-listing-aggregator/src/query"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
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

func getListingDescription(url string) string {
	html, err := getListingPage(url)
	if err != nil {
		var result string
		return result
	}

	document, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println(err)
		var result string
		return result
	}

	document.Find("script").Remove()
	tables := document.Find("body > table > tbody > tr > td > table")

	return goquery.NewDocumentFromNode(tables.Get(1)).Text()
}

func formatPublishingDateString(text string) string {
	dateLayout := "02.01.06"

	if text == "днес" {
		return time.Now().Format(dateLayout)
	} else if text == "вчера" {
		return time.Now().AddDate(0, 0, -1).Format(dateLayout)
	} else {
		return text
	}
}

// deconstructInfoTagsElement splits the text of the element and returns the first token as the first
// return value and the remaining elements as the second return value.
func deconstructInfoTagsElement(element *goquery.Selection) (string, []string) {
	text := element.Text()
	tokens := strings.Split(text, "; ")

	//trim tokens
	for i := 0; i < len(tokens); i++ {
		tokens[i] = strings.TrimSpace(tokens[i])
	}

	return tokens[0], tokens[1:]
}

// deconstructListingAnchor returns the element text, full link derived from the href value, and the id
// parameter of said link.
func deconstructListingAnchor(element *goquery.Selection) (string, string, string) {
	text := element.Text()
	href := element.AttrOr("href", "")
	link := domain + href
	idParameter := strings.Replace(href, "job/", "", -1)

	return text, link, idParameter
}
