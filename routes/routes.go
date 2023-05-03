package routes

import (
	"PongPedia/controllers"
	m "PongPedia/middleware"
	"PongPedia/models"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	// Inisialisasi Echo
	e := echo.New()

	// Middleware
	m.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	// Validator
	e.Validator = &models.CustomValidator{Validator: validator.New()}

	// Route Login and Register and Logout
	e.POST("/register", controllers.RegisterControllers)
	e.POST("/login", controllers.LoginController)
	e.POST("/logout", controllers.LogoutControllers)

	// Route Home
	e.GET("/", controllers.GetAllTurnamentControllers)

	// Route User Profile
	e.GET("/profile", controllers.GetUserControllers, m.IsLoggedIn)
	e.GET("/profile/turnament", controllers.MyTurnamentControllers, m.IsLoggedIn)
	e.PUT("/:username/profile-update", controllers.UpdateUserControllers, m.IsLoggedIn)
	e.DELETE("/:username/profile-delete", controllers.DeteleUserControllers, m.IsLoggedIn)

	// Routue Player
	// e.PUT("/player", controllers.CreatePlayersControllers, m.IsLoggedIn)
	e.PUT("/player-edit", controllers.CreatePlayersControllers, m.IsLoggedIn)

	//Routes Turnament
	e.GET("/turnaments/:id", controllers.GetTurnamentDetailControllers)
	e.POST("/turnaments", controllers.CreateTurnamentControllers)
	e.POST("/turnaments/:id/register", controllers.RegisterTurnamentControllers, m.IsLoggedIn)

	return e
}
