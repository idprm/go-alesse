package controller

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

type PharmacyProcessRequest struct {
	ChatID       uint64 `query:"chat_id" validate:"required" json:"chat_id"`
	ApothecaryID uint   `query:"apothecary_id" validate:"required" json:"apothecary_id"`
	IsProcess    bool   `query:"is_process" json:"is_process"`
}

type PharmacySentRequest struct {
	ChatID       uint64 `query:"chat_id" validate:"required" json:"chat_id"`
	ApothecaryID uint   `query:"apothecary_id" validate:"required" json:"apothecary_id"`
	IsSent       bool   `query:"is_sent" json:"is_sent"`
}

func ValidatePharmacyProcess(req PharmacyProcessRequest) []*ErrorResponse {
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

func ValidatePharmacySent(req PharmacySentRequest) []*ErrorResponse {
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

func GetAllPharmacy(c *fiber.Ctx) error {
	var pharmacies []model.Pharmacy
	database.Datasource.DB().Find(&pharmacies)
	return c.Status(fiber.StatusOK).JSON(pharmacies)
}

func SaveProcessPharmacy(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(PharmacyProcessRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidatePharmacyProcess(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var pharmacy model.Pharmacy
	existPharmacy := database.Datasource.DB().Where("chat_id", req.ChatID).First(&pharmacy)

	if existPharmacy.RowsAffected == 0 {
		pharmacy := model.Pharmacy{
			ChatID:       req.ChatID,
			ApothecaryID: req.ApothecaryID,
			IsProcess:    true,
			ProcessAt:    time.Now(),
		}
		database.Datasource.DB().Create(&pharmacy)
	} else {
		pharmacy.IsProcess = true
		pharmacy.ProcessAt = time.Now()
		database.Datasource.DB().Save(&pharmacy)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
		"data":    pharmacy,
	})
}

func SaveSentPharmacy(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(PharmacySentRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidatePharmacySent(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var pharmacy model.Pharmacy
	existPharmacy := database.Datasource.DB().Where("chat_id", req.ChatID).First(&pharmacy)

	if existPharmacy.RowsAffected == 0 {
		pharmacy := model.Pharmacy{
			ChatID:       req.ChatID,
			ApothecaryID: req.ApothecaryID,
			IsSent:       true,
			SentAt:       time.Now(),
		}
		database.Datasource.DB().Create(&pharmacy)
	} else {
		pharmacy.IsSent = true
		pharmacy.SentAt = time.Now()
		database.Datasource.DB().Save(&pharmacy)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
		"data":    pharmacy,
	})
}
