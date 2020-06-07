package listing

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
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

func scanRows(rows *sql.Rows) Listing {
	var result Listing

	rows.Scan(
		&result.ID, &result.ExternalID, &result.Link, &result.Name,
		&result.Company, &result.WorkSchedule, &result.Location,
		&result.PostingDate, &result.Description, pq.Array(&result.Keywords),
		&result.QueryID, &result.SourceName)

	return result
}

func (r *repositoryImplementation) getListingByID(id uint64) (Listing, error) {
	var result Listing

	query := `
		SELECT 
			listing_id, 
			external_id, 
			link, 
			name, 
			company,
			work_schedule, 
			location, 
			posting_date,
			description,
			keywords,
			query_id,
			source_name
		FROM
			listing
		WHERE
			listing_id = $1`

	rows, err := r.database.Query(query, id)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows.Close()

	if rows.Next() {
		result = scanRows(rows)
	}

	return result, nil
}

func (r *repositoryImplementation) getListingsByQueryIDAndSourceName(
	queryID uint64, sourceName string) ([]Listing, error) {

	var results []Listing

	query := `
	SELECT 
		listing_id, 
		external_id, 
		link, 
		name, 
		company,
		work_schedule, 
		location, 
		posting_date,
		description,
		keywords,
		query_id,
		source_name
	FROM
		listing
	WHERE
		query_id = $1
	AND
		source_name = $2`

	rows, err := r.database.Query(query, queryID, sourceName)
	if err != nil {
		log.Println(err)
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		result := scanRows(rows)
		results = append(results, result)
	}

	return results, nil
}

func (r *repositoryImplementation) getListingsByQueryID(queryID uint64) ([]Listing, error) {
	var results []Listing

	query := `
	SELECT 
		listing_id, 
		external_id, 
		link, 
		name, 
		company,
		work_schedule, 
		location, 
		posting_date,
		description,
		keywords,
		query_id,
		source_name
	FROM
		listing
	WHERE
		query_id = $1`

	rows, err := r.database.Query(query, queryID)
	if err != nil {
		log.Println(err)
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		result := scanRows(rows)
		results = append(results, result)
	}

	return results, nil
}

func (r *repositoryImplementation) getLatestListingsBySourceName(
	sourceName string, limit uint16) ([]Listing, error) {

	var results []Listing

	query := `
	SELECT 
		listing_id, 
		external_id, 
		link, 
		name, 
		company,
		work_schedule, 
		location, 
		posting_date,
		description,
		keywords,
		query_id,
		source_name
	FROM
		listing
	WHERE
		source_name = $1
	LIMIT $2`

	rows, err := r.database.Query(query, sourceName, limit)
	if err != nil {
		log.Println(err)
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		result := scanRows(rows)
		results = append(results, result)
	}

	return results, nil
}
