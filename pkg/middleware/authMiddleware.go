package middleware

import (
	"net/http"
	"restaurant-management/pkg/helpers"
	"strings"

	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		splitToken := strings.Split(authToken, "Bearer ")
		if len(splitToken) != 2 {
			return c.JSON(http.StatusUnauthorized, "User unauthothorized")
		}
		reqToken := splitToken[1]
		if reqToken == "" {
			return c.JSON(http.StatusInternalServerError, "failed to getting your token")
		}
		ok, err := helpers.ValidateToken(reqToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		if !ok {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return next(c)
	}
}

