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
	userController := controllers.NewUserController(userUsecase, userRepository)

	playerRepository := database.NewPlayerRespository(db)
	playerUsecase := usecase.NewPlayerUsecase(playerRepository, userRepository)
	playerController := controllers.NewPlayerController(playerUsecase, playerRepository)

	authRepository := database.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepository, userRepository)
	authController := controllers.NewAuthController(authUsecase, authRepository, userUsecase)
	participationRepository := database.NewParticipationRepository(db)

	turnamentRepository := database.NewTurnamentRepository(db)
	turnamentUsecase := usecase.NewTurnamentUsecase(turnamentRepository, playerRepository, userRepository, participationRepository)
	turnamentController := controllers.NewTurnamentControllers(turnamentUsecase, turnamentRepository)

	// Validator
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.POST("/login", authController.LoginUserController)
	e.POST("/register", authController.RegisterUserController)
	// e.POST("/logout", authController.LogoutUserController)

	pf := e.Group("/profile", m.IsLoggedIn)
	pf.GET("", userController.GetUserController)
	pf.PUT("", userController.UpdateUserController)
	pf.DELETE("", userController.DeleteUserController)

	pp := e.Group("/profile/player", m.IsLoggedIn)
	pp.GET("", playerController.GetPlayerController)
	pp.POST("", playerController.CreatePlayerController)
	pp.PUT("", playerController.UpdatePlayerController)

	tt := e.Group("/tournament", m.IsLoggedIn)
	tt.GET("", turnamentController.GetTurnamentController)
	tt.GET("/:id", turnamentController.GetTurnamentDetailController)
	tt.POST("", turnamentController.CreateTurnamentController)
	tt.POST("/register", turnamentController.RegisterTurnamentController)

}
