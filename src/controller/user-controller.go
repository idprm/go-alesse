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

func GetTransactionByUser(c *fiber.Ctx) error {
	var user model.User
	database.Datasource.DB().Where("msisdn", c.Params("msisdn")).Where("is_active", true).First(&user)

	var transactions []model.Transaction
	database.Datasource.DB().Where("user_id", user.ID).Order("created_at DESC").Limit(15).Find(&transactions)
	return c.Status(fiber.StatusOK).JSON(&transactions)
}
