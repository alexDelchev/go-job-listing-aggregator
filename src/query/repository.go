package query

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
)

type repository interface {
	getQueryByID(id uint64) (Query, error)

	getAllQueries() ([]Query, error)

	getActiveQueries() ([]Query, error)

	getInactiveQueries() ([]Query, error)

	insertQuery(query Query) (uint64, error)

	updateQuery(query Query) (bool, error)
}

type repositoryImplementation struct {
	database *sql.DB
}

func scanRows(rows *sql.Rows) Query {
	var result Query

	rows.Scan(
		&result.ID, pq.Array(&result.Keywords), &result.Location, &result.Active,
		&result.CreationDate)

	return result
}

func (r *repositoryImplementation) getQueryByID(id uint64) (Query, error) {
	var result Query
	query := `
		SELECT
			query_id,
			keywords,
			location,
			active,
			creation_date
		FROM
			query
		WHERE
			query_id = $1`

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

func (r *repositoryImplementation) getAllQueries() ([]Query, error) {
	var results []Query

	query := `
		SELECT
			query_id,
			keywords,
			location,
			active,
			creation_date
		FROM
			query`

	rows, err := r.database.Query(query)
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
