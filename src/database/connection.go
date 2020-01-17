package database

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
	"github.com/theodo/scalab/config"
)

var Db *sql.DB

func Connect() {
	dbCfg := config.Cfg.Database
	connector, err := pq.NewConnector(dbCfg.Url)
	if err != nil {
		log.Panic("Could not connect to database", err)
	}

	Db = sql.OpenDB(connector)
	//defer db.Close()
}
