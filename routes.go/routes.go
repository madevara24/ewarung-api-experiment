package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	// conf, err := config.LoadConfig()
	// if err != nil {
	// 	log.Fatal("Error load config on routes.go: ", err)
	// }
	// secret := conf.JWTSecret
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// DEFAULT ROUTE
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "WHAT? HOW? WHY?")
	})

	return e
}
