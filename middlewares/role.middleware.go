package middlewares

import (
	"ewarung-api-experiment/jwt"
	"ewarung-api-experiment/models"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func CheckAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)
		claims, verified := jwt.ExtractClaims(token)
		if verified && claims["role"] == "ADMIN" {
			return next(c)
		}

		var res models.Response

		res.Status = "Error"
		res.Message = "Only admin can access this endpoint"
		res.Data = nil

		c.Error(c.JSON(http.StatusUnauthorized, res))

		return nil
	}
}
