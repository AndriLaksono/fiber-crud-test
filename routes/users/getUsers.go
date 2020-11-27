package routes

import (
	database "fiber-go-crud/database/user"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	users, err := database.GetUsers()
	if err != nil {
		c.Send([]byte(err.Error()))
		c.Status(500)
		return nil
	}
	if err := c.JSON(users); err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return nil
	}

	response := fiber.Map{
		"status":  200,
		"message": "ready",
		"data":    users,
	}

	return c.Status(200).JSON(response)
}
