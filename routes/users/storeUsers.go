package routes

import (
	database "fiber-go-crud/database/user"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func StoreUser(c *fiber.Ctx) error {
	req := new(UserReq)
	if err := c.BodyParser(req); err != nil {
		fmt.Println("woy")
		return err
	}

	err := database.StoreUser(req.Username, req.Password)
	if err != nil {
		c.Send([]byte(err.Error()))
		c.Status(500)
		return nil
	}

	response := fiber.Map{
		"status":  200,
		"message": "INSERT SUCCESS",
	}

	return c.Status(200).JSON(response)
}
