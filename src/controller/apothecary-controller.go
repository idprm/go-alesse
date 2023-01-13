package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

func GetAllApothecary(c *fiber.Ctx) error {
	var apothecaries []model.Apothecary
	database.Datasource.DB().Where("is_active", true).Preload("Healthcenter").Find(&apothecaries)
	return c.Status(fiber.StatusOK).JSON(apothecaries)
}

func GetApothecaryByHealthCenter(c *fiber.Ctx) error {
	var apothecaries []model.Apothecary
	healthcenterId := c.Params("healthcenter")
	database.Datasource.DB().Where("is_active", true).Where("healthcenter_id", healthcenterId).Preload("Healthcenter").Find(&apothecaries)
	return c.Status(fiber.StatusOK).JSON(apothecaries)
}

func GetApothecaryByChannel(c *fiber.Ctx) error {
	var apothecaries []model.Apothecary
	channel := c.Params("channel")
	database.Datasource.DB().Raw("SELECT a.* FROM apothecaries a LEFT JOIN chats b ON b.healthcenter_id = a.healthcenter_id WHERE b.channel_url = ?", channel).Scan(&apothecaries)
	return c.Status(fiber.StatusOK).JSON(apothecaries)
}

func GetApothecary(c *fiber.Ctx) error {
	var apothecary model.Apothecary
	database.Datasource.DB().Where("phone", c.Params("phone")).Preload("Healthcenter").First(&apothecary)
	return c.Status(fiber.StatusOK).JSON(apothecary)
}
