package routes

import (
	"PongPedia/constants"
	"PongPedia/controllers"
	m "PongPedia/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	m.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	u := e.Group("/user")
	u.GET("", controllers.GetUserControllers, echojwt.JWT([]byte(constants.SCREAT_JWT)))
	u.POST("/register", controllers.CreateUserControllers)
	u.POST("/login", controllers.LoginUserController)

	p := e.Group("/player")
	p.GET("", controllers.GetPlayerControllers)
	p.POST("/register/turnament", controllers.CreatePlayerControllers)

	h := e.Group("/home")
	h.GET("", controllers.HomeController, echojwt.JWT([]byte(constants.SCREAT_JWT)))

	return e
}
