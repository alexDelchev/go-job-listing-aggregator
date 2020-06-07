package query

import (
	"database/sql"

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
