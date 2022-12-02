package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

func GetUserByMsisdn(c *fiber.Ctx) error {
	var user model.User
	database.Datasource.DB().Where("msisdn", c.Params("msisdn")).Where("is_active", true).First(&user)
	return c.Status(fiber.StatusOK).JSON(&user)
}
