package main

import (
	"github.com/labstack/echo/v4"
)

func main() {

	// Inisialisasi Echo dari package routes
	e := echo.New()

	// Run Server
	e.Logger.Fatal(e.Start(":8080"))
}
