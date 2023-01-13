package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

func GetAllDoctor(c *fiber.Ctx) error {
	var doctors []model.Doctor
	database.Datasource.DB().Where("is_active", true).Preload("Healthcenter").Order("end desc").Find(&doctors)
	return c.Status(fiber.StatusOK).JSON(doctors)
}

func GetDoctorByHealthCenter(c *fiber.Ctx) error {
	var doctors []model.Doctor
	healthcenterId := c.Params("healthcenter")
	database.Datasource.DB().Where("is_active", true).Where("healthcenter_id", healthcenterId).Preload("Healthcenter").Order("end desc").Find(&doctors)
	return c.Status(fiber.StatusOK).JSON(doctors)
}

func GetDoctorByChannel(c *fiber.Ctx) error {
	var doctors []model.Doctor
	channel := c.Params("channel")
	database.Datasource.DB().Raw("SELECT a.* FROM doctors a LEFT JOIN chats b ON b.healthcenter_id = a.healthcenter_id WHERE b.channel_url = ?", channel).Scan(&doctors)
	return c.Status(fiber.StatusOK).JSON(doctors)
}

func GetDoctor(c *fiber.Ctx) error {
	username := c.Params("username")
	var doctor model.Doctor
	database.Datasource.DB().Where("username", username).Preload("Healthcenter").First(&doctor)
	return c.Status(fiber.StatusOK).JSON(doctor)
}

func GetAllSpecialist(c *fiber.Ctx) error {
	var specialists []model.Specialist
	database.Datasource.DB().Where("is_active", true).Order("end desc").Find(&specialists)
	return c.Status(fiber.StatusOK).JSON(specialists)
}

func GetSpecialistByHealthCenter(c *fiber.Ctx) error {
	var specialists []model.Specialist
	healthcenterId := c.Params("healthcenter")
	database.Datasource.DB().Where("is_active", true).Where("healthcenter_id", healthcenterId).Order("end desc").Find(&specialists)
	return c.Status(fiber.StatusOK).JSON(specialists)
}

func GetSpecialist(c *fiber.Ctx) error {
	username := c.Params("username")
	var specialis model.Specialist
	database.Datasource.DB().Where("username", username).First(&specialis)
	return c.Status(fiber.StatusOK).JSON(specialis)
}
