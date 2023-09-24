package routes

import (
	"app/handler"
	"github.com/gofiber/fiber/v2"
)

func fleetRoutes(v1 fiber.Router) {
	v1.Get("/fleet", handler.GetFleets)
	v1.Post("/fleet", handler.CreateFleet)
	v1.Delete("/fleet/:id", handler.DeleteFleet)
	v1.Put("/fleet/:id", handler.UpdateFleet)
}
