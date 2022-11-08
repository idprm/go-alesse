package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

func GetAllDisease(c *fiber.Ctx) error {
	var diasease []model.Disease
	database.Datasource.DB().Order("name asc").Find(&diasease)
	return c.Status(fiber.StatusOK).JSON(&diasease)
}

func GetAllMedicine(c *fiber.Ctx) error {
	var medicines []model.Medicine
	database.Datasource.DB().Where("is_active", true).Find(&medicines)
	return c.Status(fiber.StatusOK).JSON(&medicines)
}

func GetMedicine(c *fiber.Ctx) error {
	var medicine model.Medicine
	database.Datasource.DB().Where("id", c.Params("id")).Where("is_active", true).First(&medicine)
	return c.Status(fiber.StatusOK).JSON(&medicine)
}

func GetAllHealthCenter(c *fiber.Ctx) error {
	var healthcenters []model.Healthcenter
	database.Datasource.DB().Order("name asc").Find(&healthcenters)
	return c.Status(fiber.StatusOK).JSON(healthcenters)
}

func GetHealthCenter(c *fiber.Ctx) error {
	code := c.Params("code")
	var healthcenter model.Healthcenter
	database.Datasource.DB().Where("code", code).First(&healthcenter)
	return c.Status(fiber.StatusOK).JSON(healthcenter)
}
