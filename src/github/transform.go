package github

import (
	"fmt"
	"go-job-listing-aggregator/src/listing"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SourceName defines github as the data source
const SourceName string = "github"

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

func transformToListingModel(queryID uint64, apiModel *jobListingAPIModel) listing.Listing {
	keywords := strings.Split(apiModel.Title, " ")
	descriptionText := extractTextFromHTML(apiModel.Description)

	return listing.Listing{
		ExternalID:   apiModel.ID,
		Link:         apiModel.URL,
		Name:         apiModel.Title,
		WorkSchedule: apiModel.Type,
		Company:      apiModel.Company,
		Location:     apiModel.Location,
		PostingDate:  apiModel.CreatedAt,
		Description:  descriptionText,
		Keywords:     keywords,
		QueryID:      queryID,
		SourceName:   SourceName}
}
