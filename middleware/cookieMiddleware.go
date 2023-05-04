package middleware

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func CreateCookie(c echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "JWTCookie"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
}

func DeleteCookie(c echo.Context) error {
	cookie, err := c.Cookie("JWTCookie")
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-1 * time.Hour)
	c.SetCookie(cookie)

	return nil
}
