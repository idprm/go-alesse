package controller

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

type HomecareRequest struct {
	ChatID             uint64    `query:"chat_id" validate:"required" json:"chat_id"`
	EarlyDiagnosis     string    `query:"early_diagnosis" validate:"required" json:"early_diagnosis"`
	Reason             string    `query:"reason" validate:"required" json:"reason"`
	VisitAt            time.Time `query:"visit_at" validate:"required" json:"visit_at"`
	Slug               string    `query:"slug" json:"slug"`
	Treatment          string    `query:"treatment" json:"treatment"`
	FinalDiagnosis     string    `query:"final_diagnosis" json:"final_diagnosis"`
	DrugAdministration string    `query:"drug_administration" json:"drug_administration"`
}

func ValidateHomecare(req HomecareRequest) []*ErrorResponse {
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

func GetAllHomecare(c *fiber.Ctx) error {
	var homecares []model.Homecare
	database.Datasource.DB().Find(&homecares)
	return c.Status(fiber.StatusOK).JSON(homecares)
}

func GetHomecare(c *fiber.Ctx) error {
	var homecare model.Homecare
	database.Datasource.DB().First(&homecare)
	return c.Status(fiber.StatusOK).JSON(homecare)
}

func SaveHomecare(c *fiber.Ctx) error {

	c.Accepts("application/json")

	req := new(HomecareRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateHomecare(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var homecare model.Homecare
	isExist := database.Datasource.DB().Where("chat_id", req.ChatID).First(&homecare)

	if isExist.RowsAffected == 0 {
		database.Datasource.DB().Create(
			&model.Homecare{
				ChatID:             req.ChatID,
				EarlyDiagnosis:     req.EarlyDiagnosis,
				Reason:             req.Reason,
				VisitAt:            req.VisitAt,
				Slug:               "",
				Treatment:          req.Treatment,
				FinalDiagnosis:     req.FinalDiagnosis,
				DrugAdministration: req.DrugAdministration,
			},
		)
	} else {
		homecare.EarlyDiagnosis = req.EarlyDiagnosis
		homecare.Reason = req.Reason
		homecare.VisitAt = req.VisitAt
		homecare.Slug = ""
		homecare.Treatment = req.Treatment
		homecare.FinalDiagnosis = req.FinalDiagnosis
		homecare.DrugAdministration = req.DrugAdministration
		database.Datasource.DB().Save(&homecare)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
		"data":    homecare,
	})
}
