package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

func GetAllDriver(c *fiber.Ctx) error {
	var drivers []model.Driver
	database.Datasource.DB().Order("name asc").Find(&drivers)
	return c.Status(fiber.StatusOK).JSON(&drivers)
}

func GetDriver(c *fiber.Ctx) error {
	var driver model.Driver
	database.Datasource.DB().Where("phone", c.Params("phone")).First(&driver)
	return c.Status(fiber.StatusOK).JSON(driver)
}
