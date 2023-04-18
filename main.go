package main

import (
	"PongPedia/config"
	"PongPedia/routes"
)

func main() {
	config.InitDB()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}
