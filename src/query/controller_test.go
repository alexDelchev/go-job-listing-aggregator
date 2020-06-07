package query

import (
	"errors"
)

type serviceMock struct{}

const errorQueryID uint64 = 99

func (s serviceMock) GetQueryByID(id uint64) (Query, error) {
	if id == 0 {
		return Query{}, errors.New("Sample error")
	}

	return Query{ID: id}, nil
}

func (s serviceMock) GetActiveQueries() ([]Query, error) {
	return []Query{}, errors.New("Sample error")
}

func (s serviceMock) GetInactiveQueries() ([]Query, error) {
	return []Query{}, errors.New("Sample error")
}

func (s serviceMock) GetAllQueries() ([]Query, error) {
	return []Query{}, errors.New("Sample error")
}

func (s serviceMock) CreateQuery(query Query) (Query, error) {
	if query.ID == errorQueryID {
		return query, errors.New("Sample error")
	}

	return query, nil
}

func (s serviceMock) UpdateQuery(query Query) (Query, error) {
	if query.ID == errorQueryID {
		return query, errors.New("Sample error")
	}

	return query, nil
}

func (s serviceMock) ActivateQuery(id uint64) (Query, error) {
	if id == errorQueryID {
		return Query{}, errors.New("Sample error")
	}

	return Query{}, nil
}

func (s serviceMock) DeactivateQuery(id uint64) (Query, error) {
	if id == errorQueryID {
		return Query{}, errors.New("Sample error")
	}

	return Query{}, nil
}
