package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

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
