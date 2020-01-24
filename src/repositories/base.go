package repositories

import (
	"github.com/theodo/scalab/src/database"
)

func getOneResult(dest []interface{}, query string, args ...interface{}) error {
	row := database.Db.QueryRow(query, args...)
	if err := row.Scan(dest...); err != nil {
		return &GetMappedResultError{err}
	}
	return nil
}
