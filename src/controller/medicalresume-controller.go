package controller

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

type MedicalResumeRequest struct {
	ChatID    uint64 `query:"chat_id" validate:"required" json:"chat_id"`
	DiseaseID uint   `query:"disease_id" json:"disease_id"`
	Diagnosis string `query:"diagnosis" json:"diagnosis"`
}

func ValidateMedicalResume(req MedicalResumeRequest) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func GetAllMedicalResume(c *fiber.Ctx) error {
	var medicalresumes []model.MedicalResume
	database.Datasource.DB().Find(&medicalresumes)
	return c.Status(fiber.StatusOK).JSON(medicalresumes)
}

func GetMedicalResume(c *fiber.Ctx) error {
	var medicalresume model.MedicalResume
	database.Datasource.DB().Where("number", c.Query("number")).First(&medicalresume)
	return c.Status(fiber.StatusOK).JSON(medicalresume)
}

func SaveMedicalResume(c *fiber.Ctx) error {

	c.Accepts("application/json")

	req := new(MedicalResumeRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateMedicalResume(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var medicalresume model.MedicalResume
	existResume := database.Datasource.DB().Where("chat_id", req.ChatID)
	if existResume.RowsAffected == 0 {
		database.Datasource.DB().Create(
			&model.MedicalResume{
				ChatID:    req.ChatID,
				Number:    "",
				DiseaseID: req.DiseaseID,
				Diagnosis: req.Diagnosis,
			},
		)
	} else {
		medicalresume.DiseaseID = req.DiseaseID
		medicalresume.Diagnosis = req.Diagnosis
		database.Datasource.DB().Save(&medicalresume)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
		"data":    medicalresume,
	})
}
