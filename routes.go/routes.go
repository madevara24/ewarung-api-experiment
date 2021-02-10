package routes

import (
	"ewarung-api-experiment/config"
	"ewarung-api-experiment/controllers"
	"ewarung-api-experiment/middlewares"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error load config on routes.go: ", err)
		panic(err)
	}
	secret := conf.JWTSecret
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// DEFAULT ROUTE
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "WHAT? HOW? WHY?")
	})

	// API ROUTES
	api := e.Group("/api/v1")

	// AUTH ROUTES
	api.POST("/login", controllers.Login)
	api.POST("/register", controllers.Register)

	// PROTECTED ROUTES (LOGGED IN USERS ONLY)
	protected := api.Group("/protected")
	protected.Use(middleware.JWT([]byte(secret)))

	// ROUTES FOR ADMIN
	adminRoutes := protected.Group("")
	adminRoutes.Use(middlewares.CheckAdmin)

	// BASIC USER ROUTES
	user := adminRoutes.Group("/user")
	user.GET("", controllers.GetAllUser)
	user.GET("/:id", controllers.GetUserById)
	user.POST("", controllers.StoreUser)
	user.PUT("/:id", controllers.UpdateUser)
	user.DELETE("/:id", controllers.DeleteUser)

	return e
}
