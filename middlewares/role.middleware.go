package middlewares

import (
	"ewarung-api-experiment/jwt"
	"ewarung-api-experiment/models"
	"net/http"

	"github.com/labstack/echo"
)

func CheckAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, verified := jwt.ExtractClaims(c.Request().Header.Get("Authorization"))
		if verified && claims["role"] == "ADMIN" {
			return next(c)
		}

		var res models.Response

		res.Message = "Only admin can access this endpoint"
		res.Data = nil

		c.Error(c.JSON(http.StatusUnauthorized, res))

		return nil
	}
}

func CheckShopOwner(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, verified := jwt.ExtractClaims(c.Request().Header.Get("Authorization"))
		if verified && claims["role"] == "SHOP_OWNER" {
			return next(c)
		}

		var res models.Response

		res.Message = "Only shop owners can access this endpoint"
		res.Data = nil

		c.Error(c.JSON(http.StatusUnauthorized, res))

		return nil
	}
}

func CheckCashier(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, verified := jwt.ExtractClaims(c.Request().Header.Get("Authorization"))
		if verified && claims["role"] == "CASHIER" {
			return next(c)
		}

		var res models.Response

		res.Message = "Only cashiers can access this endpoint"
		res.Data = nil

		c.Error(c.JSON(http.StatusUnauthorized, res))

		return nil
	}
}
