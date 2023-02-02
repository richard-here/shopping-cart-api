package main

import (
	"richard-here/haioo-api/cart-api/database"
	"richard-here/haioo-api/cart-api/handler"
	"richard-here/haioo-api/cart-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	database.Connect()
	handler.InitHandler()

	app := Setup()
	app.Listen(":8080")
}

func Setup() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	router.SetupRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	return app
}
