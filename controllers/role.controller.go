package controllers

import (
	"ewarung-api-experiment/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllRole(c echo.Context) error {
	result, err := models.GetAllRole()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreRole(c echo.Context) (err error) {
	u := new(models.Role)
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreRole(u.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateRole(c echo.Context) (err error) {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)

	u := new(models.Role)
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateRole(conv_id, u.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteRole(c echo.Context) (err error) {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)

	result, err := models.DeleteRole(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
