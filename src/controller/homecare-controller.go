package controller

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

type HomecareRequest struct {
	ChatID         uint64                    `query:"chat_id" validate:"required" json:"chat_id"`
	PainComplaints string                    `query:"pain_complaints" validate:"required" json:"pain_complaints"`
	EarlyDiagnosis string                    `query:"early_diagnosis" validate:"required" json:"early_diagnosis"`
	Reason         string                    `query:"reason" validate:"required" json:"reason"`
	VisitAt        string                    `query:"visit_at" validate:"required" json:"visit_at"`
	Slug           string                    `query:"slug" json:"slug"`
	Data           []HomecareMedicineRequest `query:"data" json:"data"`
}

type HomecareMedicineRequest struct {
	HomecareID  uint64 `query:"homecare_id" json:"homecare_id"`
	MedicineID  uint64 `query:"medicine_id" json:"medicine_id"`
	Name        string `query:"name" json:"name"`
	Qty         uint   `query:"quantity" json:"quantity"`
	Rule        string `query:"rule" json:"rule"`
	Dose        string `query:"dose" json:"dose"`
	Description string `query:"description" json:"description"`
}

type HomecareOfficerRequest struct {
	HomecareID uint64 `query:"homecare_id" validate:"required" json:"homecare_id"`
	Slug       string `query:"slug" json:"slug"`
	DoctorID   uint   `query:"doctor_id" json:"doctor_id"`
	OfficerID  uint   `query:"officer_id" json:"officer_id"`
	DriverID   uint   `query:"driver_id" json:"driver_id"`
}

type HomecareResumeRequest struct {
	HomecareID           uint64 `query:"homecare_id" validate:"required" json:"homecare_id"`
	Slug                 string `query:"slug" json:"slug"`
	PhysicaleExamination string `query:"physicale_examination" json:"physicale_examination"`
	FinalDiagnosis       string `query:"final_diagnosis" json:"final_diagnosis"`
	MedicalTreatment     string `query:"medical_treatment" json:"medical_treatment"`
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

func ValidateHomecareOfficer(req HomecareOfficerRequest) []*ErrorResponse {
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

func ValidateHomecareResume(req HomecareResumeRequest) []*ErrorResponse {
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

func GetHomecareByOfficer(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var homecareOfficer model.HomecareOfficer
	database.Datasource.DB().Joins("Homecare", database.Datasource.DB().Where(&model.Homecare{Slug: c.Params("channel_url")})).Preload("Doctor").Preload("Officer").Preload("Driver").First(&homecareOfficer)
	return c.Status(fiber.StatusOK).JSON(homecareOfficer)
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

func GetAllHomecareByMedicines(c *fiber.Ctx) error {
	var medicines []model.HomecareMedicine
	database.Datasource.DB().Joins("Homecare", database.Datasource.DB().Where(&model.Homecare{Slug: c.Params("slug")})).Preload("Medicine").Find(&medicines)
	return c.Status(fiber.StatusOK).JSON(medicines)
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
	// var notifDoctorToHomecare model.Config
	// database.Datasource.DB().Where("name", valDoctorToHomecare).First(&notifDoctorToHomecare)
	// replaceMessage := strings.Replace(notifDoctorToHomecare.Value, "@doctor", "", 1)

	// zenzivaDoctorToHomecare, err := handler.ZenzivaSendSMS(homecare.Chat.User.Msisdn, replaceMessage)
	// if err != nil {
	// 	// util.Log.WithFields(logrus.Fields{"error": err.Error()}).Error(valDoctorToHomecare)
	// 	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
	// 		"error":   true,
	// 		"message": err.Error(),
	// 	})
	// }
	// // insert to zenziva
	// database.Datasource.DB().Create(
	// 	&model.Zenziva{
	// 		Msisdn:   homecare.Chat.User.Msisdn,
	// 		Action:   valDoctorToHomecare,
	// 		Response: zenzivaDoctorToHomecare,
	// 	},
	// )

	visitAt, _ := time.Parse("2006-01-02 15:04", req.VisitAt)

	if isExist.RowsAffected == 0 {
		homecare := model.Homecare{
			ChatID:         req.ChatID,
			PainComplaints: req.PainComplaints,
			EarlyDiagnosis: req.EarlyDiagnosis,
			Reason:         req.Reason,
			VisitAt:        visitAt,
			Slug:           req.Slug,
			SubmitedAt:     time.Now(),
			IsSubmited:     true,
		}
		database.Datasource.DB().Create(&homecare)

		database.Datasource.DB().Where("homecare_id", homecare.ID).Delete(&model.HomecareMedicine{})

		for _, v := range req.Data {
			database.Datasource.DB().Create(&model.HomecareMedicine{
				HomecareID:  homecare.ID,
				MedicineID:  v.MedicineID,
				Name:        v.Name,
				Quantity:    v.Qty,
				Rule:        v.Rule,
				Dose:        v.Dose,
				Description: v.Description,
			})
		}

	} else {
		homecare.PainComplaints = req.PainComplaints
		homecare.EarlyDiagnosis = req.EarlyDiagnosis
		homecare.Reason = req.Reason
		homecare.VisitAt = visitAt
		homecare.Slug = req.Slug
		homecare.SubmitedAt = time.Now()
		homecare.IsSubmited = true
		database.Datasource.DB().Save(&homecare)

		database.Datasource.DB().Where("homecare_id", homecare.ID).Delete(&model.HomecareMedicine{})

		for _, v := range req.Data {
			database.Datasource.DB().Create(&model.HomecareMedicine{
				HomecareID:  homecare.ID,
				MedicineID:  v.MedicineID,
				Name:        v.Name,
				Quantity:    v.Qty,
				Rule:        v.Rule,
				Dose:        v.Dose,
				Description: v.Description,
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
		"data":    homecare,
	})
}

func SaveHomecareOfficer(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(HomecareOfficerRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateHomecareOfficer(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var homecare model.HomecareOfficer
	isExist := database.Datasource.DB().Where("homecare_id", req.HomecareID).First(&homecare)

	if isExist.RowsAffected == 0 {
		homecare := model.HomecareOfficer{
			HomecareID: req.HomecareID,
			Slug:       req.Slug,
			OfficerID:  req.OfficerID,
			DoctorID:   req.DoctorID,
			DriverID:   req.DriverID,
			VisitedAt:  time.Now(),
			IsVisited:  true,
		}
		database.Datasource.DB().Create(&homecare)
	} else {
		homecare.OfficerID = req.OfficerID
		homecare.DoctorID = req.DoctorID
		homecare.DriverID = req.DriverID
		homecare.VisitedAt = time.Now()
		homecare.IsVisited = true

		database.Datasource.DB().Save(&homecare)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Visited",
		"data":    homecare,
	})
}

func SaveHomecareResume(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(HomecareResumeRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateHomecareResume(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var homecare model.HomecareOfficer
	isExist := database.Datasource.DB().Where("homecare_id", req.HomecareID).First(&homecare)

	if isExist.RowsAffected > 0 {
		homecare.PhysicaleExamination = req.PhysicaleExamination
		homecare.FinalDiagnosis = req.FinalDiagnosis
		homecare.MedicalTreatment = req.MedicalTreatment
		homecare.FinishedAt = time.Now()
		homecare.IsFinished = true
		database.Datasource.DB().Save(&homecare)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
		"data":    homecare,
	})
}

func HomecarePhoto(c *fiber.Ctx) error {
	healthcenterId, _ := strconv.Atoi(c.FormValue("healthcenter_id"))
	homecareId, _ := strconv.ParseUint(c.FormValue("homecare_id"), 0, 64)
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
	imageFile := fmt.Sprintf("%s.%s", filename, fileExt)

	// save image to ./images dir
	err = c.SaveFile(file, fmt.Sprintf("./public/uploads/homecare/%s", imageFile))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   false,
			"message": "Server error",
			"data":    nil,
		})
	}

	database.Datasource.DB().Create(
		&model.HomecareUpload{
			HealthcenterID: uint(healthcenterId),
			HomecareID:     homecareId,
			Filename:       filename + "." + fileExt,
		},
	)

	// err := database.NewRedisClient().RPush().Err();
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Image uploaded successfully",
		"data":    imageFile,
	})
}
