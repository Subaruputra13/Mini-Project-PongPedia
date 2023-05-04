package middleware

import (
	"PongPedia/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Create Token Jwt
func CreateToken(userId int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["role_type"] = role
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SCREAT_JWT))
}

func ExtracTokenAdmin(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["role_type"] != "admin" {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}

	return 0, nil
}

func ExtracTokenUser(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["role_type"] != "player" {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}

	return 0, nil
}
