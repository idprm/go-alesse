package controller

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

type PrescriptionRequest struct {
	RequestType     string                        `query:"request_type" validate:"required" json:"request_type"`
	ChatID          uint64                        `query:"chat_id" validate:"required" json:"chat_id"`
	AllergyMedicine string                        `query:"allergy_medicine" json:"allergy_medicine"`
	Diagnosis       string                        `query:"diagnosis" json:"diagnosis"`
	Data            []PrescriptionMedicineRequest `query:"data" json:"data"`
}

type PrescriptionMedicineRequest struct {
	PrescriptionID uint64 `query:"prescription_id" json:"prescription_id"`
	MedicineID     uint64 `query:"medicine_id" json:"medicine_id"`
	Name           string `query:"medicine_name" json:"medicine_name"`
	Qty            uint   `query:"qty" json:"qty"`
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
	var medicines []model.PrescriptionMedicine
	database.Datasource.DB().Joins("Prescription", database.Datasource.DB().Where(&model.Prescription{Slug: c.Params("slug")})).Preload("Medicine").Find(&medicines)
	return c.Status(fiber.StatusOK).JSON(medicines)
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

		prescription := model.Prescription{
			ChatID:          req.ChatID,
			AllergyMedicine: req.AllergyMedicine,
			Diagnosis:       req.Diagnosis,
		}
		database.Datasource.DB().Create(&prescription)

		database.Datasource.DB().Where("prescription_id", prescription.ID).Delete(&model.PrescriptionMedicine{})

		for _, v := range req.Data {
			database.Datasource.DB().Create(&model.PrescriptionMedicine{
				PrescriptionID: prescription.ID,
				MedicineID:     v.MedicineID,
				Name:           v.Name,
				Quantity:       v.Qty,
			})
		}
	} else {
		prescription.AllergyMedicine = req.AllergyMedicine
		prescription.Diagnosis = req.Diagnosis
		database.Datasource.DB().Save(&prescription)

		database.Datasource.DB().Where("prescription_id", prescription.ID).Delete(&model.PrescriptionMedicine{})

		for _, v := range req.Data {
			database.Datasource.DB().Create(&model.PrescriptionMedicine{
				PrescriptionID: prescription.ID,
				MedicineID:     v.MedicineID,
				Name:           v.Name,
				Quantity:       v.Qty,
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
		"data":    prescription,
	})
}
