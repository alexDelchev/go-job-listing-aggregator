package weworkremotely

import (
	"fmt"
	"go-job-listing-aggregator/src/listing"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SourceName defines WeWorkRemotely as the data source
const SourceName string = "weworkremotely"

func extractTextFromHTML(data string) string {
	var html string

	// Package text in a html tag if not present
	if strings.Contains(data, "<html>") {
		html = data
	} else {
		html = fmt.Sprintf("<html><body>%s</body></html>", data)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println(err)
		return data
	}

	return doc.Text()
}

// Returns the (company name, position title) wich construct the title property.
func deconstructTitle(title string) (string, string) {
	tokens := strings.SplitN(title, ":", 2)

	return strings.TrimSpace(tokens[0]), strings.TrimSpace(tokens[1])
}

func transformToListingModel(queryID uint64, rssModel *jobListingRSSModel) listing.Listing {
	descriptionText := extractTextFromHTML(rssModel.Description)
	company, positionName := deconstructTitle(rssModel.Title)
	keywords := strings.Split(positionName, " ")

	return listing.Listing{
		ExternalID:  rssModel.ID,
		Link:        rssModel.Link,
		Name:        positionName,
		Company:     company,
		Keywords:    keywords,
		PostingDate: rssModel.PubDate,
		Description: descriptionText,
		QueryID:     queryID,
		SourceName:  SourceName}
}

func transformToListingModelSlice(queryID uint64, rssModels []jobListingRSSModel) []listing.Listing {
	var result []listing.Listing

	for _, rssModel := range rssModels {
		model := transformToListingModel(queryID, &rssModel)
		result = append(result, model)
	}

	return result
}
