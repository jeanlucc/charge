package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/theodo/scalab/src/security"
)

func ConfigureRoleHierarchy(h map[string][]string) {

}

func IsGranted(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			up := security.NewUserFromContextProvider()
			roles, err := up.GetRoles(c)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			if "" == role {
				return next(c)
			}
			for _, ur := range roles {
				if ur == role {
					return next(c)
				}
			}
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
	}
}
