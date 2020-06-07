// +build test

package stackoverflow

import (
	"fmt"
	"go-job-listing-aggregator/src/listing"
	"go-job-listing-aggregator/src/testutils"
	"testing"
)

func TestExtractTextFromHTML(t *testing.T) {
	const input string = "<div>Sample <span>text</span></div>"

	expectedOutput := "Sample text"

	actualOutput := extractTextFromHTML(input)

	if expectedOutput != actualOutput {
		t.Errorf("%s failed: Expected output %s, got %s",
			t.Name(), expectedOutput, actualOutput)
	}
}

func TestTransformToListingModel(t *testing.T) {
	model := jobListingRSSModel{
		ID:          "id",
		Link:        "link",
		Title:       "title",
		AuthorName:  "author_name",
		Categories:  []string{"value_1", "value_2", "value_3"},
		Description: "description"}

	const queryID uint64 = 1

	expectedListing := listing.Listing{
		ExternalID:  "id",
		Link:        "link",
		Name:        "title",
		Company:     "author_name",
		Keywords:    []string{"value_1", "value_2", "value_3"},
		Description: "description",
		QueryID:     queryID,
		SourceName:  SourceName}

	actualListing := transformToListingModel(queryID, &model)

	errorMessagePrefix := fmt.Sprintf("%s failed:", t.Name())
	testutils.CompareListings(t, expectedListing, actualListing, errorMessagePrefix)
}
