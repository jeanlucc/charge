package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/theodo/scalab/src/security"
)

func SessionAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			up := security.NewUserFromContextProvider()
			_, err := up.GetId(c)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			return next(c)
		}
	}
}
