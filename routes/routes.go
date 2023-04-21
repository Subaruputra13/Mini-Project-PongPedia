package routes

import (
	"PongPedia/controllers"
	m "PongPedia/middleware"
	"PongPedia/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// Inisialisasi Echo
	e := echo.New()

	// Middleware
	m.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	//Validator Require
	e.Validator = &models.CustomValidator{Validators: validator.New()}

	// Route Login and Register
	e.POST("/register", controllers.CreateUserControllers)
	e.POST("/login", controllers.LoginController)

	// Route Untuk user
	u := e.Group("/users")
	u.GET("", controllers.GetUserControllers)
	u.GET("/:id", controllers.GetUserByIdControllers)
	u.PUT("/:id", controllers.UpdateUserByIdControllers)
	u.DELETE("/:id", controllers.DeteleUserByIdControllers)

	return e
}
