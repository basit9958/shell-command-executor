package main

import (
	"github.com/basit9958/shell-command-executor/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a Fiber named instance
	app := fiber.New()
	// Set up the routing path for the POST method
	setupRoutes(app)

	err := app.Listen(":3000")
	// Raise the Panic if found error
	panic(err)
}

func setupRoutes(app *fiber.App) {
	app.Post("/api/cmd", routers.Execute)
}
