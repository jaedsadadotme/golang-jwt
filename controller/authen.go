package controller

import (
	database "demo-postgres/config"
	"demo-postgres/lib"
	models "demo-postgres/model"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type jwtCustomClaims struct {
	ID       uint   `json: id`
	Username string `json: username`
	jwt.StandardClaims
}

func comparePassword(password string, database_password string) string {
	if err := bcrypt.CompareHashAndPassword([]byte(database_password), []byte(password)); err != nil {
		return string("Invalid Credentials.")
	}
	return string("")
}

// Login ...
func Login(c echo.Context) error {
	type LoginRequest struct {
		Username string `json:username`
		Password string `json:password`
	}
	input := new(LoginRequest)
	if err := c.Bind(input); err != nil {
		return c.JSON(400, models.ErrorResponse{
			Error:   true,
			Message: "Please review your input",
		})
	}
	// not error
	user := new(models.User)

	datas := database.DB.Where(&models.User{Username: input.Username}).First(&user)
	if datas.RowsAffected <= 0 {
		return c.JSON(400, models.ErrorResponse{
			Error:   true,
			Message: "Invalid Credentials.",
		})
	}

	fmt.Println("user => ", user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.JSON(400, models.ErrorResponse{
			Error:   true,
			Message: "password incorrect",
		})
	}

	token := lib.GenerateJWT(user.ID, user.Username)
	if err := datas.Update("token", token).Error; err != nil {
		return c.JSON(400, models.ErrorResponse{
			Error:   true,
			Message: "Something went wrong, please try again later. 😕",
		})
	}
	fmt.Println("", datas)
	return c.JSON(200, models.Map{
		"message": "Login success",
		"datas": map[string]interface{}{
			"username": user.Username,
			"email":    user.Email,
			"token":    user.Token,
		},
	})
}

// Register ...
func Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(400, models.ErrorResponse{
			Error:   true,
			Message: "Please review your input",
		})
	}
	user.Token = lib.GenerateJWT(user.ID, user.Username)
	if err := database.DB.Create(&user).Error; err != nil {
		return c.JSON(400, models.ErrorResponse{
			Error:   true,
			Message: "Something went wrong, please try again later. 😕",
		})
	} else {
		return c.JSON(201, models.Map{
			"messages": "Create User Success",
		})
	}
}
