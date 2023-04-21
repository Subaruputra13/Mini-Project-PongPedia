package main

import (
	"PongPedia/config"
	"PongPedia/routes"
)

func main() {
	// Inisialisasi Databases
	config.Init()

	// Inisialisasi Echo dari package routess
	e := routes.New()

	// Run Servers
	e.Logger.Fatal(e.Start(":8080"))
}
