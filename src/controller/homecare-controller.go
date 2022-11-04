package controller

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/handler"
	"github.com/idprm/go-alesse/src/pkg/model"
)

type HomecareRequest struct {
	ChatID             uint64 `query:"chat_id" validate:"required" json:"chat_id"`
	EarlyDiagnosis     string `query:"early_diagnosis" validate:"required" json:"early_diagnosis"`
	Reason             string `query:"reason" validate:"required" json:"reason"`
	VisitAt            string `query:"visit_at" validate:"required" json:"visit_at"`
	Slug               string `query:"slug" json:"slug"`
	IsSubmited         bool   `query:"is_submited" json:"is_submited"`
	Treatment          string `query:"treatment" json:"treatment"`
	FinalDiagnosis     string `query:"final_diagnosis" json:"final_diagnosis"`
	DrugAdministration string `query:"drug_administration" json:"drug_administration"`
}

const (
	valDoctorToHomecare = "NOTIF_DOCTOR_TO_HOMECARE"
)

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
	c.Accepts("application/json")

	channelUrl := c.Query("channel_url")
	var homecare model.Homecare
	database.Datasource.DB().Where("slug", channelUrl).First(&homecare)
	return c.Status(fiber.StatusOK).JSON(homecare)
}

func GetHomecareByDoctor(c *fiber.Ctx) error {
	var homecare model.Homecare
	database.Datasource.DB().Where("slug", c.Params("slug")).Preload("Chat.Doctor").Preload("Chat.User").First(&homecare)
	return c.Status(fiber.StatusOK).JSON(homecare)
}

func GetHomecareAllPhoto(c *fiber.Ctx) error {
	var photos []model.Photo
	database.Datasource.DB().Where("homecare_id", c.Query("id")).Find(&photos)
	return c.Status(fiber.StatusOK).JSON(photos)
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

	// send notif
	var notifDoctorToHomecare model.Config
	database.Datasource.DB().Where("name", valDoctorToHomecare).First(&notifDoctorToHomecare)
	replaceMessage := strings.Replace(notifDoctorToHomecare.Value, "@doctor", "", 1)

	zenzivaDoctorToHomecare, err := handler.ZenzivaSendSMS(homecare.Chat.User.Msisdn, replaceMessage)
	if err != nil {
		// util.Log.WithFields(logrus.Fields{"error": err.Error()}).Error(valDoctorToHomecare)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// insert to zenziva
	database.Datasource.DB().Create(
		&model.Zenziva{
			Msisdn:   homecare.Chat.User.Msisdn,
			Action:   valDoctorToHomecare,
			Response: zenzivaDoctorToHomecare,
		},
	)

	visitAt, _ := time.Parse("2006-01-02 15:04:05", req.VisitAt)

	if isExist.RowsAffected == 0 {
		database.Datasource.DB().Create(
			&model.Homecare{
				ChatID:             req.ChatID,
				EarlyDiagnosis:     req.EarlyDiagnosis,
				Reason:             req.Reason,
				VisitAt:            visitAt,
				Slug:               req.Slug,
				IsSubmited:         req.IsSubmited,
				Treatment:          req.Treatment,
				FinalDiagnosis:     req.FinalDiagnosis,
				DrugAdministration: req.DrugAdministration,
			},
		)
	} else {
		homecare.EarlyDiagnosis = req.EarlyDiagnosis
		homecare.Reason = req.Reason
		homecare.VisitAt = visitAt
		homecare.Slug = req.Slug
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

func UploadPhoto(c *fiber.Ctx) error {
	// homecareId = c.FormValue("homecare_id")
	file, err := c.FormFile("photo")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   false,
			"message": "Server error",
			"data":    nil,
		})
	}

	// generate new uuid for image name
	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename
	fileExt := strings.Split(file.Filename, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	// save image to ./images dir
	err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   false,
			"message": "Server error",
			"data":    nil,
		})
	}

	database.Datasource.DB().Create(
		&model.Photo{
			HomecareID: 1,
			FileName:   image,
		},
	)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Image uploaded successfully",
		"data":    image,
	})
}
