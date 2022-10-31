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
