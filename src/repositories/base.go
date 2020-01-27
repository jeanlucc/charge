package repositories

import (
	"github.com/theodo/scalab/src/database"
)

type objectToMapRetriever func() (interface{}, []interface{})

func getOneResult(dest []interface{}, query string, args ...interface{}) error {
	row := database.Db.QueryRow(query, args...)
	if err := row.Scan(dest...); err != nil {
		return &GetMappedResultError{err}
	}
	return nil
}

func getResults(get objectToMapRetriever, query string, args ...interface{}) (results []interface{}, error error) {
	rows, err := database.Db.Query(query, args...)
	if err != nil {
		error = &GetMappedResultError{err}
		return
	}
	defer rows.Close()
	for rows.Next() {
		e, f := get()
		if err := rows.Scan(f...); err != nil {
			error = &GetMappedResultError{err}
			return
		}

		results = append(results, e)
	}
	if err := rows.Err(); err != nil {
		error = &GetMappedResultError{err}
		return
	}
	return
}
