package controllers

import (
	"ewarung-api-experiment/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllShop(c echo.Context) error {
	result, err := models.GetAllShop()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
func GetShopById(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.GetUserById(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func StoreShop(c echo.Context) (err error) {
	u := new(models.Shop)
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreShop(u.Name, u.Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateShop(c echo.Context) (err error) {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)

	u := new(models.Shop)
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateShop(conv_id, u.Name, u.Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteShop(c echo.Context) (err error) {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)

	result, err := models.DeleteShop(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
