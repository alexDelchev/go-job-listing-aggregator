// +build test

package weworkremotely

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

func TestDeconstructTitle(t *testing.T) {
	input := "value1:value2:value3"

	expectedFirstValue := "value1"
	expectedSecondValue := "value2:value3"

	actualFistValue, actualSecondValue := deconstructTitle(input)

	if expectedFirstValue != actualFistValue {
		t.Errorf("%s failed: Expected first value %s, got %s",
			t.Name(), expectedFirstValue, actualFistValue)
	}

	if expectedSecondValue != actualSecondValue {
		t.Errorf("%s failed: Expected sconed value %s, got %s",
			t.Name(), expectedSecondValue, actualSecondValue)
	}
}

func TestTransformToListingModel(t *testing.T) {
	model := jobListingRSSModel{
		ID:          "id",
		Link:        "link",
		Title:       "company: value1 value2",
		PubDate:     "publish_date",
		Description: "description"}

	const queryID uint64 = 1

	expectedListing := listing.Listing{
		ExternalID:  "id",
		Link:        "link",
		Name:        "value1 value2",
		Company:     "company",
		Keywords:    []string{"value1", "value2"},
		PostingDate: "publish_date",
		Description: "description",
		QueryID:     queryID,
		SourceName:  SourceName}

	actualListing := transformToListingModel(queryID, &model)

	errorMessagePrefix := fmt.Sprintf("%s failed:", t.Name())
	testutils.CompareListings(t, expectedListing, actualListing, errorMessagePrefix)
}
