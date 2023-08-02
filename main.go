package main

import (
	"PongPedia/config"
	"PongPedia/middleware"
	"PongPedia/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env")
	}

	db := config.InitDB()
	e := echo.New()

	routes.NewRoute(e, db)
	middleware.LogMiddleware(e)

	e.Logger.Fatal(e.Start(":8080"))
}
