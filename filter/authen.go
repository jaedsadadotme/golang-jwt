package filter

import (
	models "demo-postgres/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("authorization")
		claims := new(models.Claims)
		if token == "" {
			return c.JSON(403, "Unauthorize")
		} else {
			_, err := jwt.ParseWithClaims(token, claims,
				func(token *jwt.Token) (interface{}, error) {
					return []byte("demojwt"), nil
				})
			if err != nil {
				return c.JSON(403, models.ErrorResponse{
					Error:   true,
					Message: "token is not valid",
				})
			}
		}
		return next(c)
	}
}
