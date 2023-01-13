package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

func GetAllDriver(c *fiber.Ctx) error {
	var drivers []model.Driver
	database.Datasource.DB().Order("name asc").Preload("Healthcenter").Find(&drivers)
	return c.Status(fiber.StatusOK).JSON(&drivers)
}

func GetDriverByHealthCenter(c *fiber.Ctx) error {
	var drivers []model.Driver
	healthcenterId := c.Params("healthcenter")
	database.Datasource.DB().Where("healthcenter_id", healthcenterId).Order("name asc").Preload("Healthcenter").Find(&drivers)
	return c.Status(fiber.StatusOK).JSON(&drivers)
}

func GetDriverByChannel(c *fiber.Ctx) error {
	var drivers []model.Driver
	channel := c.Params("channel")
	database.Datasource.DB().Raw("SELECT a.* FROM drivers a LEFT JOIN chats b ON b.healthcenter_id = a.healthcenter_id WHERE b.channel_url = ?", channel).Scan(&drivers)
	return c.Status(fiber.StatusOK).JSON(drivers)
}

func GetDriver(c *fiber.Ctx) error {
	var driver model.Driver
	database.Datasource.DB().Where("phone", c.Params("phone")).Preload("Healthcenter").First(&driver)
	return c.Status(fiber.StatusOK).JSON(driver)
}
