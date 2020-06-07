package listing

import (
	"database/sql"
)

type repository interface {
	getListingByID(id uint64) (Listing, error)

	getListingsByQueryIDAndSourceName(queryID uint64, sourceName string) ([]Listing, error)

	getListingsByQueryID(queryID uint64) ([]Listing, error)

	getLatestListingsBySourceName(sourceName string, limit uint16) ([]Listing, error)

	getSourceNames() ([]string, error)

	listingExists(externalID string, sourceName string) (bool, error)

	insertListing(listing *Listing) (uint64, error)
}

type repositoryImplementation struct {
	database *sql.DB
}
