package controllers

import (
	"ewarung-api-experiment/jwt"
	"ewarung-api-experiment/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllUser(c echo.Context) error {
	result, err := models.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetUserById(c echo.Context) error {
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

func StoreUser(c echo.Context) (err error) {
	u := new(models.User)
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreUser(u.Username, u.Email, u.Password, u.FullName, u.IDRole)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUser(c echo.Context) (err error) {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)

	u := new(models.User)
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateUser(conv_id, u.Username, u.Email, u.Password, u.FullName, u.IDRole)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context) (err error) {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)

	result, err := models.DeleteUser(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func GetLogedInUserWithRole(c echo.Context) (err error) {
	claims, verified := jwt.ExtractClaims(c.Request().Header.Get("Authorization"))
	if !verified {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	conv_id := int(claims["id"].(float64))

	result, err := models.GetUserWithRoleById(conv_id)

	return c.JSON(http.StatusOK, result)
}
