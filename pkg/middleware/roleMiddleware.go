package middleware

import (
	"net/http"
	"restaurant-management/pkg/config"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		hToken := c.Request().Header.Get("Authorization")
		token := strings.Split(hToken, " ")[1]
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.LocalConfig.SECRETKEY), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to parse token")
		}
		if claims["UserType"] != "admin" {
			return echo.NewHTTPError(http.StatusForbidden, "Only admin can access")
		}
		return next(c)
	}
}
