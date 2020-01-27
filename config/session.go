package config

import (
	"log"
	"time"

	"github.com/antonlindstrom/pgstore"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func ConfSession(e *echo.Echo) {
	databaseCfg := GetDatabase()
	sessionCfg := GetSession()
	store, err := pgstore.NewPGStore(databaseCfg.Url, []byte(sessionCfg.Secret))
	if err != nil {
		log.Panic("Could not configure session store")
	}
	store.Cleanup(time.Minute * 5)
	store.MaxAge(sessionCfg.MaxAge)

	e.Use(session.Middleware(store))
}
