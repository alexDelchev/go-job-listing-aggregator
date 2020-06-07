// +build test

package testutils

import (
	"go-job-listing-aggregator/src/listing"
	"testing"
)

func CompareStringSlices(a []string, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for index, value := range a {
		if value != b[index] {
			return false
		}
	}

	return true
}

func CompareListings(t *testing.T, expected listing.Listing, actual listing.Listing, errorMessagePrefix string) {
	if expected.ExternalID != actual.ExternalID {
		t.Errorf("%s Exptected ExterrnalID %s, got %s", errorMessagePrefix, expected.ExternalID, actual.ExternalID)
	}

	if expected.Link != actual.Link {
		t.Errorf("%s Expected Link %s, got %s", errorMessagePrefix, expected.Link, actual.Link)
	}

	if expected.Name != actual.Name {
		t.Errorf("%s Expected Name %s, got %s", errorMessagePrefix, expected.Name, actual.Name)
	}

	if expected.WorkSchedule != actual.WorkSchedule {
		t.Errorf("%s Expected WorkSchedule %s, got %s", errorMessagePrefix, expected.WorkSchedule, actual.WorkSchedule)
	}

	if expected.Company != actual.Company {
		t.Errorf("%s Expected Comapny %s, got %s", errorMessagePrefix, expected.Company, actual.Company)
	}

	if expected.Location != actual.Location {
		t.Errorf("%s Expected Location %s, got %s", errorMessagePrefix, expected.Location, actual.Location)
	}

	if expected.PostingDate != actual.PostingDate {
		t.Errorf("%s Expected PostingDate %s, got %s", errorMessagePrefix, expected.PostingDate, actual.PostingDate)
	}

	if expected.Description != actual.Description {
		t.Errorf("%s Expected Description %s, got %s", errorMessagePrefix, expected.Description, actual.Description)
	}

	if !CompareStringSlices(expected.Keywords, actual.Keywords) {
		t.Errorf("%s Expected Keywords %+v, got %+v", errorMessagePrefix, expected.Keywords, actual.Keywords)
	}

	if expected.QueryID != actual.QueryID {
		t.Errorf("%s Expected QueryID %d, got %d", errorMessagePrefix, expected.QueryID, actual.QueryID)
	}

	if expected.SourceName != actual.SourceName {
		t.Errorf("%s Expected SourceNa,e %s, got %s", errorMessagePrefix, expected.SourceName, actual.SourceName)
	}
}
