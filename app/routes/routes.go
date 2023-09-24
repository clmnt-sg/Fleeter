package routes

import "github.com/gofiber/fiber/v2"

func FleeterRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	userRoutes(v1)

}
