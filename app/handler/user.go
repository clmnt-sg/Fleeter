package handler

import (
	"app/config"
	"app/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	config.DB.Find(&users)
	return c.Status(http.StatusOK).JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.DB.Create(&user)
	return c.Status(http.StatusCreated).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	ID := c.Params("id")
	result := config.DB.Where("ID = ?", ID).Delete(&models.User{})
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(ID)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

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
		"message": "User successfully updated.",
	})
}
