package controller

import (
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/handler"
	"github.com/idprm/go-alesse/src/pkg/model"
	"github.com/idprm/go-alesse/src/pkg/util"
)

type MedicalResumeRequest struct {
	ChatID         uint64 `query:"chat_id" validate:"required" json:"chat_id"`
	Weight         uint   `query:"weight" validate:"required" json:"weight"`
	PainComplaints string `query:"pain_complaints" json:"pain_complaints"`
	Diagnosis      string `query:"diagnosis" json:"diagnosis"`
	DiseaseID      uint   `query:"disease_id"  validate:"required" json:"disease_id"`
	Slug           string `query:"slug" json:"slug"`
	IsSubmited     bool   `query:"is_submited" json:"is_submited"`
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
	database.Datasource.DB().Where("slug", c.Params("slug")).First(&medicalresume)
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
	existResume := database.Datasource.DB().Where("chat_id", req.ChatID).First(&medicalresume)
	if existResume.RowsAffected == 0 {
		database.Datasource.DB().Create(
			&model.MedicalResume{
				ChatID:         req.ChatID,
				Weight:         req.Weight,
				PainComplaints: req.PainComplaints,
				Diagnosis:      req.Diagnosis,
				Number:         util.TimeStamp(),
				Slug:           req.Slug,
				DiseaseID:      req.DiseaseID,
				IsSubmited:     req.IsSubmited,
			},
		)

	} else {
		medicalresume.Weight = req.Weight
		medicalresume.PainComplaints = req.PainComplaints
		medicalresume.Diagnosis = req.Diagnosis
		medicalresume.Slug = req.Slug
		medicalresume.DiseaseID = req.DiseaseID
		medicalresume.IsSubmited = req.IsSubmited
		database.Datasource.DB().Save(&medicalresume)
	}

	const (
		valFeedbackToPatient = "FEEDBACK_TO_PATIENT"
	)

	var chat model.Chat
	database.Datasource.DB().Where("id", req.ChatID).Preload("Doctor").Preload("User").Preload("Healthcenter").First(&chat)

	var status model.Status
	database.Datasource.DB().Where("name", valFeedbackToPatient).First(&status)
	notifMessage := util.ContentFeedbackToPatient(status.ValueNotif, chat)
	userMessage := util.StatusFeedbackToPatient(status.ValueUser, chat)

	log.Println(notifMessage)
	log.Println(userMessage)

	zenzivaFeedbackToPatient, err := handler.ZenzivaSendSMS(chat.User.Msisdn, notifMessage)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// insert to zenziva
	database.Datasource.DB().Create(
		&model.Zenziva{
			Msisdn:   chat.User.Msisdn,
			Action:   valFeedbackToPatient,
			Response: zenzivaFeedbackToPatient,
		},
	)

	// insert to transaction
	database.Datasource.DB().Create(
		&model.Transaction{
			UserID:       chat.UserID,
			ChatID:       chat.ID,
			SystemStatus: status.ValueSystem,
			NotifStatus:  notifMessage,
			UserStatus:   userMessage,
		},
	)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
		"data":    medicalresume,
	})
}
