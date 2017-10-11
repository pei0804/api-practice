package mymiddleware

import (
	"net/http"

	"github.com/labstack/echo"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth != "auth" {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return next(c)
	}
}
