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
	playerRepository := database.NewPlayerRespository(db)
	playerUsecase := usecase.NewPlayerUsecase(playerRepository)
	authRepository := database.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepository)

	authController := controllers.NewAuthController(authUsecase, authRepository, userRepository)
	userController := controllers.NewUserController(userUsecase, userRepository)
	playerController := controllers.NewPlayerController(playerUsecase, playerRepository)

	// Validator
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.POST("/login", authController.LoginUserController)
	e.POST("/register", authController.RegisterUserController)

	pf := e.Group("/profile", m.IsLoggedIn)
	pf.GET("", userController.GetUserController)
	pf.PUT("", userController.UpdateUserController)
	pf.DELETE("", userController.DeleteUserController)

	pp := e.Group("/profile/player", m.IsLoggedIn)
	pp.GET("", playerController.GetPlayerController)
	pp.POST("", playerController.CreatePlayerController)
	pp.PUT("", playerController.UpdatePlayerController)

}
