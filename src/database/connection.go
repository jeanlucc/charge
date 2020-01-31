package database

import (
	"log"

	"github.com/go-pg/pg/v9"
	"github.com/theodo/scalab/config"
)

var db *pg.DB

func Db() pg.DB {
	return *db
}

func Close() {
	db.Close()
}

func Connect() {
	dbCfg := config.GetDatabase()
	options, err := pg.ParseURL(dbCfg.Url)
	if err != nil {
		log.Panic("Could not parse database connection url", err)
	}
	db = pg.Connect(options)
}
