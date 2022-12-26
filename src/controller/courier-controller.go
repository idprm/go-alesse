package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

func GetAllCourier(c *fiber.Ctx) error {
	var couriers []model.Courier
	database.Datasource.DB().Order("name asc").Preload("Healthcenter").Find(&couriers)
	return c.Status(fiber.StatusOK).JSON(&couriers)
}

func GetCourierByHealthCenter(c *fiber.Ctx) error {
	var couriers []model.Courier
	healthcenterId := c.Params("healthcenter")
	database.Datasource.DB().Where("healthcenter_id", healthcenterId).Order("name asc").Preload("Healthcenter").Find(&couriers)
	return c.Status(fiber.StatusOK).JSON(&couriers)
}

func GetCourier(c *fiber.Ctx) error {
	var courier model.Courier
	database.Datasource.DB().Where("phone", c.Params("phone")).Preload("Healthcenter").First(&courier)
	return c.Status(fiber.StatusOK).JSON(courier)
}
