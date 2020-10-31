package controller

import (
	database "demo-postgres/config"
	models "demo-postgres/model"

	"github.com/labstack/echo/v4"
)

func ListUser(c echo.Context) error {
	var user []models.User
	database.DB.Select([]string{"username", "email", "created_at"}).Find(&user)
	return c.JSON(200, user)
}
