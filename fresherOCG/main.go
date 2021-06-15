package main

import (
	"github.com/gleo08/fresherOCG/database"
	"github.com/gleo08/fresherOCG/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.SetUp(app)

	app.Listen(":8080")
}
