package router

import (
	"richard-here/haioo-api/cart-api/handler"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/cart")

	v1.Get("/", handler.CartHandler.GetAllProductsInCartHandler)
	v1.Post("/", handler.CartHandler.AddProductToCartHandler)
	v1.Delete("/:code", handler.CartHandler.RemoveProductFromCartHandler)
}
