package handler

import (
	"app/config"
	"app/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetFleets(c *fiber.Ctx) error {
	var users []models.Fleet
	config.DB.Find(&users)
	return c.Status(http.StatusOK).JSON(users)
}

func CreateFleet(c *fiber.Ctx) error {
	user := new(models.Fleet)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.DB.Create(&user)
	return c.Status(http.StatusCreated).JSON(user)
}

func DeleteFleet(c *fiber.Ctx) error {
	ID := c.Params("id")
	result := config.DB.Where("ID = ?", ID).Delete(&models.Fleet{})
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(ID)
}

func UpdateFleet(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.Fleet

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	result := config.DB.Model(user).Where("ID = ?", id).Updates(&user)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Fleet successfully updated.",
	})
}
