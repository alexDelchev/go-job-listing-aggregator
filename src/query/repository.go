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

func newRepository(database *sql.DB) repository {
	return &repositoryImplementation{database: database}
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

func (r *repositoryImplementation) getActiveQueries() ([]Query, error) {
	var results []Query

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
			active = true`

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

func (r *repositoryImplementation) getInactiveQueries() ([]Query, error) {
	var results []Query

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
			active = false`

	rows, err := r.database.Query(query)
	if err != nil {
		log.Println(err)
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		var result Query
		rows.Scan(
			&result.ID, pq.Array(&result.Keywords), &result.Location, &result.Active,
			&result.CreationDate)

		results = append(results, result)
	}

	return results, nil
}

func (r *repositoryImplementation) insertQuery(query Query) (uint64, error) {
	var result uint64

	statement := `
		INSERT INTO query(
			keywords,
			location,
			active
		) VALUES (
			$1, $2, $3
		) RETURNING query_id`

	rows, err := r.database.Query(
		statement,
		pq.Array(query.Keywords), query.Location, query.Active)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&result)
	}

	return result, nil
}

func (r *repositoryImplementation) updateQuery(query Query) (bool, error) {
	statement := `
		UPDATE 
			query
		SET
			keywords = $1,
			location = $2,
			active = $3
		WHERE
			query_id = $4`

	_, err := r.database.Exec(
		statement, pq.Array(query.Keywords), query.Location, query.Active, query.ID)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}
