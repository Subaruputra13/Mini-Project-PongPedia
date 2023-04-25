package routes

import (
	"PongPedia/controllers"
	m "PongPedia/middleware"
	"PongPedia/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	// Inisialisasi Echo
	e := echo.New()

	// Middleware
	m.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	// Validator Require
	e.Validator = &models.CustomValidator{Validators: validator.New()}

	// Route Login and Register
	e.POST("/register", controllers.CreateUserControllers)
	e.POST("/login", controllers.LoginController)

	// Route Untuk user
	u := e.Group("/users")
	u.GET("", controllers.GetUserControllers, m.IsLoggedIn, m.IsAdmin)
	u.GET("/:id", controllers.GetUserByIdControllers, m.IsLoggedIn, m.IsAdmin)
	u.PUT("/:id", controllers.UpdateUserByIdControllers, m.IsLoggedIn, m.IsAdmin)
	u.DELETE("/:id", controllers.DeteleUserByIdControllers, m.IsLoggedIn, m.IsAdmin)

	// Route untuk Player
	p := e.Group("/players")
	p.GET("", controllers.GetPlayersControllers, m.IsLoggedIn)
	p.POST("", controllers.CreatePlayersControllers, m.IsLoggedIn)

	// Route untuk Turnament
	t := e.Group("/turnaments")
	t.GET("", controllers.GetAllTurnamentControllers, m.IsLoggedIn)
	t.POST("", controllers.CreateTurnamentControllers, m.IsLoggedIn, m.IsAdmin)

	return e
}
