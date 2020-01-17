package repositories

import (
	"log"

	"github.com/theodo/scalab/src/database"
)

func getOneResult(dest []interface{}, query string, args ...interface{}) {
	row := database.Db.QueryRow(query, args...)
	if err := row.Scan(dest...); err != nil {
		log.Panic("Could not retrieve entity: ", err)
	}
	return
}
