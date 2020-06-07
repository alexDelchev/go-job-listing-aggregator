package listing

import "testing"

type repositoryMock struct {
	methodCalls map[string]int
}

const nonExistingExternalID string = "DOES_NOT_EXIST"

func (r repositoryMock) getListingByID(id uint64) (Listing, error) {
	return Listing{ID: id}, nil
}

func (r repositoryMock) getListingsByQueryIDAndSourceName(queryID uint64, sourceName string) ([]Listing, error) {
	return []Listing{{QueryID: queryID, SourceName: sourceName}, {QueryID: queryID, SourceName: sourceName}}, nil
}

func (r repositoryMock) getListingsByQueryID(queryID uint64) ([]Listing, error) {
	return []Listing{{QueryID: queryID}, {QueryID: queryID}}, nil
}

func (r repositoryMock) getLatestListingsBySourceName(sourceName string, limit uint16) ([]Listing, error) {
	return []Listing{{SourceName: sourceName}, {SourceName: sourceName}}, nil
}

func (r repositoryMock) getSourceNames() ([]string, error) {
	return []string{"value1", "value2", "value3"}, nil
}

// Returns false for externalID == "DOES_NOT_EXIST"
func (r repositoryMock) listingExists(externalID string, sourceName string) (bool, error) {
	return externalID != nonExistingExternalID, nil
}

func (r repositoryMock) insertListing(listing *Listing) (uint64, error) {
	r.methodCalls["insertListing"]++
	return 1, nil
}

func TestCreateListings(t *testing.T) {
	mock := repositoryMock{methodCalls: make(map[string]int)}

	service := NewService(mock)

	service.CreateListings(
		[]Listing{{ExternalID: nonExistingExternalID}, {ExternalID: "EXISTS"}, {ExternalID: nonExistingExternalID}})

	expectedCalls := 2
	actualCalls := mock.methodCalls["insertListing"]

	if expectedCalls != actualCalls {
		t.Errorf("%s failed: Expected %d calls of repository.insertListing, got %d",
			t.Name(), expectedCalls, actualCalls)
	}
}
