package stackoverflow

import (
	"fmt"
	"go-job-listing-aggregator/src/listing"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SourceName defines stackoverflow as the data source
const SourceName string = "stackoverflow"

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

func transformToListingModel(queryID uint64, rssModel *jobListingRSSModel) listing.Listing {
	descriptionText := extractTextFromHTML(rssModel.Description)

	return listing.Listing{
		ExternalID:  rssModel.ID,
		Link:        rssModel.Link,
		Name:        rssModel.Title,
		Company:     rssModel.AuthorName,
		Keywords:    rssModel.Categories,
		PostingDate: rssModel.PublishingDate,
		Description: descriptionText,
		QueryID:     queryID,
		SourceName:  SourceName}
}
