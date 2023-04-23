package routes

import (
	"PongPedia/constants"
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

	//Validator Require
	e.Validator = &models.CustomValidator{Validators: validator.New()}

	// Route Login and Register
	e.POST("/register", controllers.CreateUserControllers)
	e.POST("/login", controllers.LoginController)

	// Route Untuk user
	u := e.Group("/users")
	u.GET("", controllers.GetUserControllers, mid.JWT([]byte(constants.SCREAT_JWT)), m.IsAdmin)
	u.GET("/:id", controllers.GetUserByIdControllers, mid.JWT([]byte(constants.SCREAT_JWT)), m.IsAdmin)
	u.PUT("/:id", controllers.UpdateUserByIdControllers, mid.JWT([]byte(constants.SCREAT_JWT)), m.IsAdmin)
	u.DELETE("/:id", controllers.DeteleUserByIdControllers, mid.JWT([]byte(constants.SCREAT_JWT)), m.IsAdmin)

	return e
}
