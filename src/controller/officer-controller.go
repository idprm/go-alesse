package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

func GetAllOfficer(c *fiber.Ctx) error {

	var officers []model.Officer
	database.Datasource.DB().Where("is_active", true).Preload("Healthcenter").Order("id desc").Find(&officers)

	return c.Status(fiber.StatusOK).JSON(officers)
}

func GetOfficer(c *fiber.Ctx) error {

	var officer model.Officer
	database.Datasource.DB().Where("phone", c.Params("phone")).Preload("Healthcenter").First(&officer)

	return c.Status(fiber.StatusOK).JSON(officer)
}
