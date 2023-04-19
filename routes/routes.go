package routes

import (
	"PongPedia/controllers"
	m "PongPedia/middleware"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// Inisialisasi Echo
	e := echo.New()

	// Middleware
	m.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	// Route
	e.POST("/register", controllers.CreateUserControllers)
	e.POST("/login", controllers.LoginController)

	// Route Untuk user
	u := e.Group("/user")
	u.GET("", controllers.GetUserControllers)
	u.GET("/:role", controllers.GetUserByRoleControllers)
	u.PUT("/:id", controllers.UpdateUserByIdControllers)
	u.DELETE("/:id", controllers.DeteleUserByIdControllers)

	// Route Untuk player
	p := e.Group("/player")
	p.GET("", controllers.GetPlayerControllers)
	p.POST("", controllers.CreatePlayerControllers)

	t := e.Group("/turnament")
	t.GET("", controllers.GetTurnamentControllers)
	t.POST("", controllers.CreateTurnamentControllers)

	//testing branch

	return e
}
