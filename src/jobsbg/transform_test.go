// +build test

package jobsbg

import (
	"go-job-listing-aggregator/src/listing"
	"go-job-listing-aggregator/src/testutils"
	"testing"
)

func TestTransformToListingModel(t *testing.T) {
	model := jobsBGListing{
		ID:             "id",
		Company:        "company",
		Link:           "link",
		Title:          "title",
		PublishingDate: "publishing_date",
		Description:    "description",
		Location:       "location",
		Keywords:       []string{"value_1", "value_2", "value_3"}}

	var queryID uint64 = 1

	expected := listing.Listing{
		ExternalID:  "id",
		Link:        "link",
		Name:        "title",
		Company:     "company",
		Location:    "location",
		PostingDate: "publishing_date",
		Description: "description",
		Keywords:    []string{"value_1", "value_2", "value_3"},
		QueryID:     queryID,
		SourceName:  SourceName}

	actual := transformToListingModel(queryID, &model)

	const errorMessagePrefix string = "TestTransformToListingModel failed:"

	testutils.CompareListings(t, expected, actual, errorMessagePrefix)
}
