package routes

import (
	"PongPedia/constants"
	"PongPedia/controllers"
	m "PongPedia/middleware"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

var IsLoggedIn = mid.JWTWithConfig(mid.JWTConfig{
	SigningMethod: "HS256",
	SigningKey:    []byte(constants.SCREAT_JWT),
	TokenLookup:   "cookie:JWTCookie",
})

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		IsAdmin := claims["role_type"].(string)
		if IsAdmin == "PLAYER" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func New() *echo.Echo {
	// Inisialisasi Echo
	e := echo.New()

	// Middleware
	m.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	// Route Login and Register
	e.POST("/register", controllers.RegisterControllers)
	e.POST("/login", controllers.LoginController)

	// // Route Untuk User
	// u := e.Group("/users")
	// u.GET("", controllers.GetUserControllers, IsLoggedIn)
	// u.GET("/:role", controllers.GetUserByRoleControllers, IsLoggedIn, IsAdmin)
	// u.PUT("/:id", controllers.UpdateUserByIdControllers, IsLoggedIn)
	// u.DELETE("/:id", controllers.DeteleUserByIdControllers, IsLoggedIn)

	// // Route untuk Player
	// p := e.Group("/players")
	// p.GET("", controllers.GetPlayersControllers, IsLoggedIn)
	// p.POST("", controllers.CreatePlayersControllers, IsLoggedIn)

	// // Route untuk Turnament
	// t := e.Group("/turnaments")
	// t.POST("", controllers.CreateTurnamentControllers, IsLoggedIn, IsAdmin)
	// t.POST("/:id/register", controllers.RegisterTurnamentControllers, IsLoggedIn)

	// // Route untuk Match
	// m := e.Group("/matchs")
	// m.GET("", controllers.GetAllMatchControllers)
	// m.GET("/:id/result", controllers.GetMatchByIdControllers)
	// m.POST("", controllers.CreateMatchControllers, IsLoggedIn, IsAdmin)
	// m.PUT("/:id", controllers.UpdateMatchControllers, IsLoggedIn, IsAdmin)

	return e
}
