package config

import (
	"time"
	"log"

	"github.com/antonlindstrom/pgstore"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/session"
)

func Session(e *echo.Echo) {
	store, err := pgstore.NewPGStore(Cfg.Database.Url, []byte(Cfg.Session.Secret))
	if err != nil {
		log.Panic("Could not configure session store")
	}
	store.Cleanup(time.Minute * 5)
	store.MaxAge(Cfg.Session.MaxAge)

	e.Use(session.Middleware(store))
}
