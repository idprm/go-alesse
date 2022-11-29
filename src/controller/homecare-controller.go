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
	"github.com/idprm/go-alesse/src/pkg/handler"
	"github.com/idprm/go-alesse/src/pkg/model"
	"github.com/idprm/go-alesse/src/pkg/util"
)

type HomecareRequest struct {
	ChatID         uint64                    `query:"chat_id" validate:"required" json:"chat_id"`
	PainComplaints string                    `query:"pain_complaints" validate:"required" json:"pain_complaints"`
	EarlyDiagnosis string                    `query:"early_diagnosis" validate:"required" json:"early_diagnosis"`
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
	HomecareID           uint64                    `query:"homecare_id" validate:"required" json:"homecare_id"`
	Slug                 string                    `query:"slug" json:"slug"`
	PhysicaleExamination string                    `query:"physicale_examination" json:"physicale_examination"`
	FinalDiagnosis       string                    `query:"final_diagnosis" json:"final_diagnosis"`
	MedicalTreatment     string                    `query:"medical_treatment" json:"medical_treatment"`
	BloodPressure        string                    `query:"blood_pressure" json:"blood_pressure"`
	Weight               uint32                    `query:"weight" json:"weight"`
	Height               uint32                    `query:"height" json:"height"`
	Data                 []HomecareMedicineRequest `query:"data" json:"data"`
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

	var homecare model.Homecare
	database.Datasource.DB().Where("slug", c.Params("slug")).Preload("Chat").Preload("Chat.Doctor").Preload("Chat.User").First(&homecare)
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
	isExist := database.Datasource.DB().Where("chat_id", req.ChatID).Preload("Chat.User").First(&homecare)

	visitAt, _ := time.Parse("2006-01-02 15:04", req.VisitAt)

	if isExist.RowsAffected == 0 {
		homecare := model.Homecare{
			ChatID:         req.ChatID,
			PainComplaints: req.PainComplaints,
			EarlyDiagnosis: req.EarlyDiagnosis,
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

		var hc model.Homecare
		database.Datasource.DB().Where("chat_id", req.ChatID).Preload("Chat.User").Preload("Chat.Doctor").Preload("Chat.Healthcenter").First(&hc)

		var officer model.Officer
		database.Datasource.DB().Where("healthcenter_id", hc.Chat.HealthcenterID).First(&officer)

		const (
			valDoctorToHomecare = "NOTIF_DOCTOR_TO_HOMECARE"
			valMessageToUser    = "NOTIF_MESSAGE_USER"
		)
		var confDoctorToHomecare model.Config
		database.Datasource.DB().Where("name", valDoctorToHomecare).First(&confDoctorToHomecare)
		replaceMessageDoctorToHomecare := util.ContentDoctorToHomecare(confDoctorToHomecare.Value, hc)

		log.Println(replaceMessageDoctorToHomecare)

		var confMessageToUser model.Config
		database.Datasource.DB().Where("name", valMessageToUser).First(&confMessageToUser)
		replaceMessageToUser := util.ContentNotifToUser(confMessageToUser.Value, hc)

		log.Println(replaceMessageToUser)

		zenzivaDoctorToHomecare, err := handler.ZenzivaSendSMS(officer.Phone, replaceMessageDoctorToHomecare)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		// insert to zenziva
		database.Datasource.DB().Create(
			&model.Zenziva{
				Msisdn:   officer.Phone,
				Action:   valDoctorToHomecare,
				Response: zenzivaDoctorToHomecare,
			},
		)

		zenzivaMessageToUser, err := handler.ZenzivaSendSMS(hc.Chat.User.Msisdn, replaceMessageToUser)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		// insert to zenziva
		database.Datasource.DB().Create(
			&model.Zenziva{
				Msisdn:   hc.Chat.User.Msisdn,
				Action:   valMessageToUser,
				Response: zenzivaMessageToUser,
			},
		)

		// insert to notif
		database.Datasource.DB().Create(
			&model.Notif{
				UserID:  hc.Chat.UserID,
				Content: "",
			},
		)

		// insert to notif
		database.Datasource.DB().Create(
			&model.Notif{
				UserID:  hc.Chat.UserID,
				Content: "",
			},
		)

	} else {
		homecare.PainComplaints = req.PainComplaints
		homecare.EarlyDiagnosis = req.EarlyDiagnosis
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

	var homecareOfficer model.HomecareOfficer
	isExist := database.Datasource.DB().Where("homecare_id", req.HomecareID).First(&homecareOfficer)

	if isExist.RowsAffected == 0 {
		homecareOfficer := model.HomecareOfficer{
			HomecareID: req.HomecareID,
			Slug:       req.Slug,
			OfficerID:  req.OfficerID,
			DoctorID:   req.DoctorID,
			DriverID:   req.DriverID,
			VisitedAt:  time.Now(),
			IsVisited:  true,
		}
		database.Datasource.DB().Create(&homecareOfficer)

	} else {
		homecareOfficer.OfficerID = req.OfficerID
		homecareOfficer.DoctorID = req.DoctorID
		homecareOfficer.DriverID = req.DriverID
		homecareOfficer.VisitedAt = time.Now()
		homecareOfficer.IsVisited = true

		database.Datasource.DB().Save(&homecareOfficer)
	}

	var hc model.Homecare
	database.Datasource.DB().Where("id", req.HomecareID).Preload("Chat.User").Preload("Chat.Doctor").Preload("Chat.Healthcenter").First(&hc)

	const (
		valHomecareToPatientProgress = "NOTIF_HOMECARE_TO_PATIENT_PROGRESS"
	)
	var conf model.Config
	database.Datasource.DB().Where("name", valHomecareToPatientProgress).First(&conf)
	replaceMessage := util.ContentHomecareToPatientProgress(conf.Value, hc)
	log.Println(replaceMessage)

	zenzivaHomecareToPatientProgress, err := handler.ZenzivaSendSMS(hc.Chat.User.Msisdn, replaceMessage)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// insert to zenziva
	database.Datasource.DB().Create(
		&model.Zenziva{
			Msisdn:   hc.Chat.User.Msisdn,
			Action:   valHomecareToPatientProgress,
			Response: zenzivaHomecareToPatientProgress,
		},
	)

	// insert to notif
	database.Datasource.DB().Create(
		&model.Notif{
			UserID:  hc.Chat.UserID,
			Content: "",
		},
	)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Visited",
		"data":    homecareOfficer,
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

	var homecareOfficer model.HomecareOfficer
	isExist := database.Datasource.DB().Where("homecare_id", req.HomecareID).First(&homecareOfficer)

	if isExist.RowsAffected > 0 {
		homecareOfficer.PhysicaleExamination = req.PhysicaleExamination
		homecareOfficer.FinalDiagnosis = req.FinalDiagnosis
		homecareOfficer.MedicalTreatment = req.MedicalTreatment
		homecareOfficer.BloodPressure = req.BloodPressure
		homecareOfficer.Height = req.Height
		homecareOfficer.Weight = req.Weight
		homecareOfficer.FinishedAt = time.Now()
		homecareOfficer.IsFinished = true
		database.Datasource.DB().Save(&homecareOfficer)
	}

	database.Datasource.DB().Where("homecare_id", req.HomecareID).Delete(&model.HomecareMedicine{})

	for _, v := range req.Data {
		database.Datasource.DB().Create(&model.HomecareMedicine{
			HomecareID:  req.HomecareID,
			MedicineID:  v.MedicineID,
			Name:        v.Name,
			Quantity:    v.Qty,
			Rule:        v.Rule,
			Dose:        v.Dose,
			Description: v.Description,
		})
	}

	var hc model.Homecare
	database.Datasource.DB().Where("id", req.HomecareID).Preload("Chat").Preload("Chat.User").Preload("Chat.Doctor").Preload("Chat.Healthcenter").First(&hc)

	var superadmin model.SuperAdmin
	database.Datasource.DB().First(&superadmin)

	const (
		valHomecareToPatientDone  = "NOTIF_HOMECARE_TO_PATIENT_DONE"
		valHomecareToHealthOffice = "NOTIF_HOMECARE_TO_HEALTHOFFICE"
	)
	var confHomecareToPatientDone model.Config
	database.Datasource.DB().Where("name", valHomecareToPatientDone).First(&confHomecareToPatientDone)
	replaceMessageHomecareToPatientDone := util.ContentHomecareToPatientDone(confHomecareToPatientDone.Value, hc)

	var confHomecareToHealthOffice model.Config
	database.Datasource.DB().Where("name", valHomecareToHealthOffice).First(&confHomecareToHealthOffice)
	replaceMessageHomecareToHealthOffice := util.ContentHomecareToHealthoffice(confHomecareToHealthOffice.Value, hc)

	zenzivaHomecareToPatientDone, err := handler.ZenzivaSendSMS(hc.Chat.User.Msisdn, replaceMessageHomecareToPatientDone)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	zenzivaHomecareToHealthOffice, err := handler.ZenzivaSendSMS(superadmin.Phone, replaceMessageHomecareToHealthOffice)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// insert to zenziva
	database.Datasource.DB().Create(
		&model.Zenziva{
			Msisdn:   hc.Chat.User.Msisdn,
			Action:   valHomecareToPatientDone,
			Response: zenzivaHomecareToPatientDone,
		},
	)

	// insert to zenziva
	database.Datasource.DB().Create(
		&model.Zenziva{
			Msisdn:   superadmin.Phone,
			Action:   valHomecareToHealthOffice,
			Response: zenzivaHomecareToHealthOffice,
		},
	)
	// insert to notif
	database.Datasource.DB().Create(
		&model.Notif{
			UserID:  hc.Chat.UserID,
			Content: "",
		},
	)

	// insert to notif
	database.Datasource.DB().Create(
		&model.Notif{
			UserID:  hc.Chat.UserID,
			Content: "",
		},
	)

	// insert to notif
	database.Datasource.DB().Create(
		&model.Notif{
			UserID:  hc.Chat.UserID,
			Content: "",
		},
	)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
		"data":    homecareOfficer,
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
