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

func GetOfficerByHealthCenter(c *fiber.Ctx) error {
	var officers []model.Officer
	healthcenterId := c.Params("healthcenter")
	database.Datasource.DB().Where("is_active", true).Where("healthcenter_id", healthcenterId).Preload("Healthcenter").Order("id desc").Find(&officers)
	return c.Status(fiber.StatusOK).JSON(officers)
}

func GetOfficerByChannel(c *fiber.Ctx) error {
	var officers []model.Officer
	channel := c.Params("channel")
	database.Datasource.DB().Raw("SELECT a.* FROM apothecaries a LEFT JOIN chats b ON b.healthcenter_id = a.healthcenter_id WHERE b.channel_url = ?", channel).Scan(&officers)
	return c.Status(fiber.StatusOK).JSON(officers)
}

func GetOfficer(c *fiber.Ctx) error {
	var officer model.Officer
	database.Datasource.DB().Where("phone", c.Params("phone")).Preload("Healthcenter").First(&officer)
	return c.Status(fiber.StatusOK).JSON(officer)
}
