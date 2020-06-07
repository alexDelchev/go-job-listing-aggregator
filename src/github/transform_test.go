// +build test

package github

import (
	"go-job-listing-aggregator/src/listing"
	"go-job-listing-aggregator/src/testutils"
	"testing"
)

func TestTransformToListingModel(t *testing.T) {
	apiModel := jobListingAPIModel{
		ID:          "id",
		Type:        "type",
		URL:         "URL",
		CreatedAt:   "created_at",
		Company:     "company",
		CompanyURL:  "company_url",
		Location:    "location",
		Title:       "value1 value2 value3",
		Description: "<p>description</p>",
		HowToApply:  "how_to_apply",
		CompanyLogo: "company_logo",
	}

	var queryID uint64 = 1

	expected := listing.Listing{
		ExternalID:   "id",
		Link:         "URL",
		Name:         "value1 value2 value3",
		WorkSchedule: "type",
		Company:      "company",
		Location:     "location",
		PostingDate:  "created_at",
		Description:  "description",
		Keywords:     []string{"value1", "value2", "value3"},
		QueryID:      queryID,
		SourceName:   SourceName}

	actual := transformToListingModel(queryID, &apiModel)

	const errorMessagePrefix string = "TestTransformToListingModel failed:"

	testutils.CompareListings(t, expected, actual, errorMessagePrefix)
}
