package controller

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

type PrescriptionRequest struct {
	ChatID          uint64 `query:"chat_id" validate:"required" json:"chat_id"`
	AllergyMedicine string `query:"allergy_medicine" json:"allergy_medicine"`
	Diagnosis       string `query:"diagnosis" json:"diagnosis"`
}

type PrescriptionMedicineRequest struct {
	PrescriptionID uint64 `query:"prescription_id" validate:"required" json:"prescription_id"`
	MedicineID     uint   `query:"medicine_id" validate:"required" json:"medicine_id"`
}

func ValidatePrescription(req PrescriptionRequest) []*ErrorResponse {
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

func GetAllPrescription(c *fiber.Ctx) error {
	var prescriptions []model.Prescription
	database.Datasource.DB().Find(&prescriptions)
	return c.Status(fiber.StatusOK).JSON(prescriptions)
}

func GetPrescription(c *fiber.Ctx) error {
	var prescription model.Prescription
	database.Datasource.DB().First(&prescription)
	return c.Status(fiber.StatusOK).JSON(prescription)
}

func GetAllPrescriptionMedicine(c *fiber.Ctx) error {
	var prescriptionmedicines model.PrescriptionMedicine
	database.Datasource.DB().Where("prescription_id", c.Query("prescription_id")).Find(&prescriptionmedicines)
	return c.Status(fiber.StatusOK).JSON(prescriptionmedicines)
}

func SavePrescription(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(PrescriptionRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidatePrescription(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var prescription model.Prescription
	existPrescription := database.Datasource.DB().Where("chat_id", req.ChatID).First(&prescription)

	if existPrescription.RowsAffected == 0 {
		database.Datasource.DB().Create(
			&model.Prescription{
				ChatID:          req.ChatID,
				AllergyMedicine: req.AllergyMedicine,
				Diagnosis:       req.Diagnosis,
			},
		)
	} else {
		prescription.AllergyMedicine = req.AllergyMedicine
		prescription.Diagnosis = req.Diagnosis
		database.Datasource.DB().Save(&prescription)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
		"data":    prescription,
	})
}
