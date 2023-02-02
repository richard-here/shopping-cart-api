package handler

import (
	"richard-here/haioo-api/cart-api/database"
	"richard-here/haioo-api/cart-api/model"
	"richard-here/haioo-api/cart-api/repository"
	"strings"

	valid "richard-here/haioo-api/cart-api/validator"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	Repo repository.Repo
}

var CartHandler Handler

func InitHandler() {
	db := database.DB
	repo := repository.CreateRepository(db)
	CartHandler = Handler{Repo: repo}
}

func (h *Handler) AddProductToCartHandler(c *fiber.Ctx) error {
	p := new(model.Product)

	err := c.BodyParser(p)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	errs := valid.ValidateProduct(*p)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": errs,
		})
	}

	p.Name = strings.ToLower(p.Name)
	existing, err := h.Repo.VerifyProductExists(p.Name)
	if existing == nil {
		err = h.Repo.AddProductToCart(p)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":  "success",
			"message": "Product was added to cart",
			"data":    p,
		})
	} else {
		p.Quantity = p.Quantity + existing.Quantity
		p.Code = existing.Code
		p.DeletedAt = gorm.DeletedAt{}
		err = h.Repo.UpdateProductInCart(p)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Product was added to cart",
			"data":    p,
		})
	}
}

func (h *Handler) GetAllProductsInCartHandler(c *fiber.Ctx) error {
	ps, err := h.Repo.GetProductsInCart()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   ps,
	})
}

func (h *Handler) RemoveProductFromCartHandler(c *fiber.Ctx) error {
	code := c.Params("code")
	err := h.Repo.DeleteProductFromCart(code)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Product removed from cart",
	})
}
