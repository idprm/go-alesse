package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

func GetAllDoctor(c *fiber.Ctx) error {

	var doctors []model.Doctor
	database.Datasource.DB().Where("is_specialist", false).Order("end desc").Find(&doctors)

	return c.Status(fiber.StatusOK).JSON(doctors)
}

func GetDoctor(c *fiber.Ctx) error {
	username := c.Params("username")

	var doctor model.Doctor
	database.Datasource.DB().Where("username", username).First(&doctor)

	return c.Status(fiber.StatusOK).JSON(doctor)
}

func GetAllDoctorSpecialist(c *fiber.Ctx) error {

	var doctors []model.Doctor
	database.Datasource.DB().Where("is_specialist", true).Order("name").Find(&doctors)

	return c.Status(fiber.StatusOK).JSON(doctors)
}
