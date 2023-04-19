package middleware

import (
	"PongPedia/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int, username, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["username"] = username
	claims["role_type"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SCREAT_JWT))
}
