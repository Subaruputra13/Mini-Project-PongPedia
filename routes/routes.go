package routes

import (
	"PongPedia/controllers"
	m "PongPedia/middleware"
	"PongPedia/repository/database"
	"PongPedia/usecase"
	utils "PongPedia/util"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	// Middleware
	e.Pre(mid.RemoveTrailingSlash())
	e.Use(mid.CORS())

	userRepository := database.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userUsecase, userRepository)

	playerRepository := database.NewPlayerRespository(db)
	playerUsecase := usecase.NewPlayerUsecase(playerRepository, userRepository)
	playerController := controllers.NewPlayerController(playerUsecase, playerRepository)

	authUsecase := usecase.NewAuthUsecase(userRepository)
	authController := controllers.NewAuthController(authUsecase, userUsecase)
	participationRepository := database.NewParticipationRepository(db)

	turnamentRepository := database.NewTurnamentRepository(db)
	turnamentUsecase := usecase.NewTurnamentUsecase(turnamentRepository, playerRepository, userRepository, participationRepository)
	turnamentController := controllers.NewTurnamentControllers(turnamentUsecase, turnamentRepository)

	matchRepository := database.NewMatchRepository(db)
	matchUsecase := usecase.NewMatchUsecase(matchRepository, participationRepository)
	matchController := controllers.NewMatchController(matchUsecase, matchRepository)

	adminUsecase := usecase.NewDashboardUsecase(userRepository, turnamentRepository, matchRepository, playerRepository)
	adminController := controllers.NewAdminControllers(adminUsecase, matchUsecase, playerUsecase, turnamentUsecase)

	// Validator
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	e.POST("/login", authController.LoginUserController)
	e.POST("/register", authController.RegisterUserController)
	e.GET("/player", adminController.GetAllPlayersController)

	// Admin Routes
	a := e.Group("/admin", m.IsLoggedIn)
	a.GET("/dashboard", adminController.DashboardAdminController)
	a.GET("/user", adminController.GetAllUserController)
	a.POST("/match", adminController.CreateMatchController)
	a.POST("/turnament", adminController.CreateTurnamentController)
	a.PUT("/match/:id", adminController.UpdateMatchController)
	a.PUT("/turnament/:id", adminController.UpdateTurnamentController)

	// User Routes
	pf := e.Group("/profile", m.IsLoggedIn)
	pf.GET("", userController.GetUserController)
	pf.PUT("", userController.UpdateUserController)
	pf.DELETE("", userController.DeleteUserController)

	// User Player Routes
	pp := e.Group("/profile/player", m.IsLoggedIn)
	pp.GET("", playerController.GetPlayerByUserIdController)
	pp.PATCH("", playerController.UpdatePlayerController)

	// Turnament Routes
	tt := e.Group("/tournament")
	tt.GET("", turnamentController.GetTurnamentController)
	tt.GET("/:id", turnamentController.GetTurnamentDetailController)
	tt.POST("/register", turnamentController.RegisterTurnamentController, m.IsLoggedIn)

	// Match Routes
	mm := e.Group("/match")
	mm.GET("", matchController.GetMatchController)
	mm.GET("/:id", matchController.GetMatchByIdController)

}
