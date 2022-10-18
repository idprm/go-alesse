package controller

import "github.com/gofiber/fiber/v2"

func FrontHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Welcome to a-lesse.com"})
}

func AuthHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Welcome to a-lesse.com"})
}
