package routes

import (
	"PongPedia/controllers"
	m "PongPedia/middleware"
	"PongPedia/repository/database"
	"PongPedia/usecase"
	"PongPedia/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {

	// Middleware
	m.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	userRepository := database.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	authController := controllers.NewUserController(userUsecase, userRepository)

	// Validator
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.POST("/login", authController.LoginUserController)
	e.POST("/register", authController.RegisterUserController)

}
