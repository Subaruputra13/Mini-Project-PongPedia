package middleware

import (
	"PongPedia/constants"
	"time"

	mid "github.com/labstack/echo/middleware"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var IsLoggedIn = mid.JWTWithConfig(mid.JWTConfig{
	SigningMethod: "HS256",
	SigningKey:    []byte(constants.SCREAT_JWT),
	TokenLookup:   "cookie:JWTCookie",
	AuthScheme:    "user",
})

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

func Auth(c echo.Context) int {
	cookie, _ := c.Cookie("JWTCookie")
	token, _ := jwt.Parse(cookie.Value, nil)
	claims := token.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	return userId

}

func IsAdmin(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["role_type"] != "admin" {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}

	return 0, nil
}

func IsUser(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["role_type"] != "player" {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}

	return 0, nil
}
