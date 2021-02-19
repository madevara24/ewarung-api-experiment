package controllers

import (
	"ewarung-api-experiment/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllUserShop(c echo.Context) error {
	result, err := models.GetAllUserShop()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
