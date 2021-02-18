package controllers

import (
	"ewarung-api-experiment/jwt"
	"ewarung-api-experiment/models"
	"ewarung-api-experiment/password"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	var res models.Response
	cred := new(models.Credentials)
	if err := c.Bind(cred); err != nil {
		return err
	}

	result, err := models.GetUserByUsername(cred.Username)

	if err != nil {
		fmt.Println(err)
		return err
	}

	isPass, err := password.CheckPasswordHash(cred.Password, result.Password)

	if err != nil || !isPass {
		fmt.Println(err)
		res.Message = "Wrong password"
		res.Data = nil
		return c.JSON(http.StatusBadRequest, res)
	}

	token, err := jwt.CreateToken(result.ID, result.Username, result.Role)

	if err != nil {
		fmt.Println(err)
		return err
	}

	res.Message = "Success login"
	res.Data = token
	return c.JSON(http.StatusOK, res)
}

func Register(c echo.Context) (err error) {
	u := new(models.User)
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreUser(u.Username, u.Email, u.Password, u.FullName, 2)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
