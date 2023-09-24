package main

import (
	"app/config"
	"app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New(config.FiberConfig())
	errGorm := config.Connect()
	if errGorm != nil {
		return
	}
	print(errGorm)
	app.Use(logger.New())
	app.Use(cors.New())
	routes.FleeterRoutes(app)

	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	errFiber := app.Listen(":3000")
	if errFiber != nil {
		return
	}
}
