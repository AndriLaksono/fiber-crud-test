package api

import (
	routesUser "fiber-go-crud/routes/users"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Init() {

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	// config bisa disetting berdasarkan kebutuhan
	configcors := cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}

	app.Use(cors.New(configcors))

	// Routes
	app.Get("/api/users", routesUser.GetAllUsers)
	app.Post("/api/users", routesUser.StoreUser)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("ERROR WOY " + err.Error())
	}
}
