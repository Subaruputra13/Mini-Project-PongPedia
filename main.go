package main

import (
	"PongPedia/config"
	"PongPedia/routes"
)

func main() {
	// Inisialisasi Database
	config.Init()

	// Inisialisasi Echo dari package routes
	e := routes.New()

	// Run Server
	e.Logger.Fatal(e.Start(":8080"))
}
