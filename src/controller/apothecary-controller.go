package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

func GetAllApothecary(c *fiber.Ctx) error {
	var apothecaries []model.Apothecary
	database.Datasource.DB().Where("is_active", true).Find(&apothecaries)
	return c.Status(fiber.StatusOK).JSON(apothecaries)
}

func GetApothecary(c *fiber.Ctx) error {
	var apothecary model.Apothecary
	database.Datasource.DB().Where("phone", c.Params("phone")).First(&apothecary)
	return c.Status(fiber.StatusOK).JSON(apothecary)
}
