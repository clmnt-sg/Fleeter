package routes

import (
	"app/handler"
	"github.com/gofiber/fiber/v2"
)

func userRoutes(v1 fiber.Router) {
	v1.Get("/user", handler.GetUsers)
	v1.Post("/user", handler.CreateUser)
	v1.Delete("/user/:id", handler.DeleteUser)
	v1.Put("/user/:id", handler.UpdateUser)
}
