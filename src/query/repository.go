package query

import (
	"database/sql"
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
