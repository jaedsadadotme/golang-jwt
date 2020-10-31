package router

import (
	"demo-postgres/controller"
	"demo-postgres/filter"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(200, "Hello World!")
}

// SetupRoutes setups all the Routes
func SetupRoutes() {
	app := echo.New()

	api := app.Group("/api")
	api.GET("/", hello)

	// authen
	Auth := api.Group("/auth")
	Auth.POST("/signin", controller.Login)
	Auth.POST("/signup", controller.Register)

	// user and require token
	USER := api.Group("/user", filter.CheckAuth)
	USER.GET("/list", controller.ListUser)

	app.Start(":3000")
}

func StartServe() {

}
