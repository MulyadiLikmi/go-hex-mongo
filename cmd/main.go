package main

import (
	"log"

	"go-hex-mongo/server"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Set up the Fiber app
	app := fiber.New()

	// Set up routes using the router
	server.SetupRoutes(app)

	// Start the Fiber server
	log.Fatal(app.Listen(":8000"))
}
