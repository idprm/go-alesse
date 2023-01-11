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
	RequestType    string                    `query:"request_type" validate:"required" json:"request_type"`
	ChatID         uint64                    `query:"chat_id" validate:"required" json:"chat_id"`
	PainComplaints string                    `query:"pain_complaints" validate:"required" json:"pain_complaints"`
	EarlyDiagnosis string                    `query:"early_diagnosis" validate:"required" json:"early_diagnosis"`
	VisitAt        string                    `query:"visit_at" validate:"required" json:"visit_at"`
	Slug           string                    `query:"slug" json:"slug"`
	Data           []HomecareMedicineRequest `query:"data" json:"data"`
	IsSoon         bool                      `query:"is_soon" json:"is_soon"`
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
	RequestType  string `query:"request_type" validate:"required" json:"request_type"`
	HomecareID   uint64 `query:"homecare_id" validate:"required" json:"homecare_id"`
	Slug         string `query:"slug" json:"slug"`
	DoctorID     uint   `query:"doctor_id" json:"doctor_id"`
	ApothecaryID uint   `query:"apothecary_id" json:"apothecary_id"`
	OfficerID    uint   `query:"officer_id" json:"officer_id"`
	DriverID     uint   `query:"driver_id" json:"driver_id"`
}

type HomecareResumeRequest struct {
	RequestType          string                    `query:"request_type" validate:"required" json:"request_type"`
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

const (
	valHomecareToPatientDone     = "HOMECARE_TO_PATIENT_DONE"
	valHomecareToHealthOffice    = "HOMECARE_TO_HEALTHOFFICE"
	valDoctorToHomecare          = "DOCTOR_TO_HOMECARE"
	valDoctorToPatientHomecare   = "DOCTOR_TO_PATIENT_HOMECARE"
	valHomecareToPatientProgress = "HOMECARE_TO_PATIENT_PROGRESS"
)

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
	var homecare model.Homecare
	database.Datasource.DB().Where("slug", c.Params("slug")).First(&homecare)

	var homecareOfficer model.HomecareOfficer
	database.Datasource.DB().Where("homecare_id", homecare.ID).Preload("Doctor").Preload("Apothecary").Preload("Officer").Preload("Driver").First(&homecareOfficer)
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
	var homecare model.Homecare
	database.Datasource.DB().Where("slug", c.Params("slug")).First(&homecare)

	var medicines []model.HomecareMedicine
	database.Datasource.DB().Where("homecare_id", homecare.ID).Preload("Medicine").Find(&medicines)
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

	// category
	var category model.Category
	database.Datasource.DB().Where("code", "homecare").First(&category)

	var hc model.Homecare
	database.Datasource.DB().Where("chat_id", req.ChatID).Preload("Chat.User").Preload("Chat.Doctor").Preload("Chat.Healthcenter").First(&hc)

	var officer model.Officer
	database.Datasource.DB().Where("healthcenter_id", hc.Chat.HealthcenterID).First(&officer)

	var statusDoctorToHomecare model.Status
	database.Datasource.DB().Where("name", valDoctorToHomecare).First(&statusDoctorToHomecare)
	notifMessageDoctorToHomecare := util.ContentDoctorToHomecare(statusDoctorToHomecare.ValueNotif, hc)
	userMessageDoctorToHomecare := util.StatusDoctorToHomecare(statusDoctorToHomecare.ValueUser, hc)
	pushMessageDoctorToHomecare := util.PushDoctorToHomecare(statusDoctorToHomecare.ValuePush, hc)

	visitAt, _ := time.Parse("2006-01-02 15:04:05", req.VisitAt)

	if isExist.RowsAffected == 0 {
		homecare := model.Homecare{
			ChatID:         req.ChatID,
			PainComplaints: req.PainComplaints,
			EarlyDiagnosis: req.EarlyDiagnosis,
			VisitAt:        visitAt.Local(),
			Slug:           req.Slug,
			SubmitedAt:     time.Now(),
			IsSoon:         req.IsSoon,
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

		log.Println(notifMessageDoctorToHomecare)
		log.Println(userMessageDoctorToHomecare)
		log.Println(pushMessageDoctorToHomecare)

		// var statusMessageToUser model.Status
		// database.Datasource.DB().Where("name", valMessageToUser).First(&statusMessageToUser)
		// notifMessageToUser := util.ContentMessageToUser(statusMessageToUser.ValueNotif, hc, officer)
		// userMessageToUser := util.StatusMessageToUser(statusMessageToUser.ValueUser, hc, officer)

		// log.Println(notifMessageToUser)
		// log.Println(userMessageToUser)

		if req.RequestType == "web" {
			zenzivaDoctorToHomecare, err := handler.ZenzivaSendSMS(officer.Phone, notifMessageDoctorToHomecare)
			if err != nil {
				return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
					"error":   true,
					"message": err.Error(),
				})
			}

			// zenzivaMessageToUser, err := handler.ZenzivaSendSMS(hc.Chat.User.Msisdn, notifMessageToUser)
			// if err != nil {
			// 	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			// 		"error":   true,
			// 		"message": err.Error(),
			// 	})
			// }

			// insert to zenziva
			database.Datasource.DB().Create(
				&model.Zenziva{
					Msisdn:   officer.Phone,
					Action:   valDoctorToHomecare,
					Response: zenzivaDoctorToHomecare,
				},
			)
		}

		// insert to transaction
		database.Datasource.DB().Create(
			&model.Transaction{
				UserID:       hc.Chat.UserID,
				ChatID:       hc.ChatID,
				SystemStatus: statusDoctorToHomecare.ValueSystem,
				NotifStatus:  notifMessageDoctorToHomecare,
				UserStatus:   userMessageDoctorToHomecare,
				PushStatus:   pushMessageDoctorToHomecare,
			},
		)

	} else {
		homecare.PainComplaints = req.PainComplaints
		homecare.EarlyDiagnosis = req.EarlyDiagnosis
		homecare.VisitAt = visitAt.Local()
		homecare.Slug = req.Slug
		homecare.SubmitedAt = time.Now()
		homecare.IsSoon = req.IsSoon
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

	var chatCategory model.ChatCategory
	isExistCategory := database.Datasource.DB().Where("chat_id", req.ChatID).First(&chatCategory)

	if isExistCategory.RowsAffected == 0 {
		database.Datasource.DB().Create(
			&model.ChatCategory{
				ChatID:     req.ChatID,
				CategoryID: category.ID,
			})
	} else {
		chatCategory.ChatID = req.ChatID
		chatCategory.CategoryID = category.ID
		database.Datasource.DB().Save(&chatCategory)
	}

	// chat closed
	var ch model.Chat
	database.Datasource.DB().Where("id", req.ChatID).First(&ch)
	ch.IsDone = true
	ch.DoneAt = time.Now()
	database.Datasource.DB().Save(&ch)

	if req.RequestType == "mobile" {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":        false,
			"message":      "Submited",
			"data":         homecare,
			"push_message": pushMessageDoctorToHomecare,
		})
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
			HomecareID:   req.HomecareID,
			Slug:         req.Slug,
			DoctorID:     req.DoctorID,
			OfficerID:    req.OfficerID,
			ApothecaryID: req.ApothecaryID,
			DriverID:     req.DriverID,
			VisitedAt:    time.Now(),
			IsVisited:    true,
		}
		database.Datasource.DB().Create(&homecareOfficer)

	} else {
		homecareOfficer.DoctorID = req.DoctorID
		homecareOfficer.OfficerID = req.OfficerID
		homecareOfficer.ApothecaryID = req.ApothecaryID
		homecareOfficer.DriverID = req.DriverID
		homecareOfficer.VisitedAt = time.Now()
		homecareOfficer.FinishedAt = time.Now()
		homecareOfficer.IsVisited = true

		database.Datasource.DB().Save(&homecareOfficer)
	}

	var hc model.Homecare
	database.Datasource.DB().Where("id", req.HomecareID).Preload("Chat.User").Preload("Chat.Doctor").Preload("Chat.Healthcenter").First(&hc)

	var officer model.Officer
	database.Datasource.DB().Where("healthcenter_id", hc.Chat.HealthcenterID).First(&officer)

	var status model.Status
	database.Datasource.DB().Where("name", valHomecareToPatientProgress).First(&status)
	notifMessageHomecareToPatientProgress := util.ContentHomecareToPatientProgress(status.ValueNotif, hc, officer)
	userMessageHomecareToPatientProgress := util.StatusHomecareToPatientProgress(status.ValueNotif, hc)
	pushMessageHomecareToPatientProgress := util.PushHomecareToPatientProgress(status.ValuePush, hc, officer)

	log.Println(notifMessageHomecareToPatientProgress)
	log.Println(userMessageHomecareToPatientProgress)
	log.Println(pushMessageHomecareToPatientProgress)

	if req.RequestType == "web" {
		zenzivaHomecareToPatientProgress, err := handler.ZenzivaSendSMS(hc.Chat.User.Msisdn, notifMessageHomecareToPatientProgress)
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

	}

	// insert to transaction
	database.Datasource.DB().Create(
		&model.Transaction{
			UserID:       hc.Chat.UserID,
			ChatID:       hc.ChatID,
			SystemStatus: status.ValueSystem,
			NotifStatus:  notifMessageHomecareToPatientProgress,
			UserStatus:   userMessageHomecareToPatientProgress,
			PushStatus:   pushMessageHomecareToPatientProgress,
		},
	)

	if req.RequestType == "mobile" {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":        false,
			"message":      "Visited",
			"data":         homecareOfficer,
			"push_message": pushMessageHomecareToPatientProgress,
		})
	}

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

	var statusHomecareToPatientDone model.Status
	database.Datasource.DB().Where("name", valHomecareToPatientDone).First(&statusHomecareToPatientDone)
	notifMessageHomecareToPatientDone := util.ContentHomecareToPatientDone(statusHomecareToPatientDone.ValueNotif, hc)
	userMessageHomecareToPatientDone := util.StatusHomecareToPatientDone(statusHomecareToPatientDone.ValueUser, hc)
	pushMessageHomecareToPatientDone := util.PushHomecareToPatientDone(statusHomecareToPatientDone.ValuePush, hc)

	log.Println(notifMessageHomecareToPatientDone)
	log.Println(userMessageHomecareToPatientDone)
	log.Println(pushMessageHomecareToPatientDone)

	var statusHomecareToHealthOffice model.Status
	database.Datasource.DB().Where("name", valHomecareToHealthOffice).First(&statusHomecareToHealthOffice)
	notifMessageHomecareToHealthOffice := util.ContentHomecareToHealthoffice(statusHomecareToHealthOffice.ValueNotif, hc)
	userMessageHomecareToHealthOffice := util.StatusHomecareToHealthoffice(statusHomecareToHealthOffice.ValueUser, hc)
	pushMessageHomecareToHealthOffice := util.PushHomecareToHealthoffice(statusHomecareToHealthOffice.ValueUser, hc)

	log.Println(notifMessageHomecareToHealthOffice)
	log.Println(userMessageHomecareToHealthOffice)
	log.Println(pushMessageHomecareToHealthOffice)

	if req.RequestType == "web" {
		zenzivaHomecareToPatientDone, err := handler.ZenzivaSendSMS(hc.Chat.User.Msisdn, notifMessageHomecareToPatientDone)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		zenzivaHomecareToHealthOffice, err := handler.ZenzivaSendSMS(superadmin.Phone, notifMessageHomecareToHealthOffice)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		// insert to zenziva
		database.Datasource.DB().Create(
			&[]model.Zenziva{{
				Msisdn:   hc.Chat.User.Msisdn,
				Action:   valHomecareToPatientDone,
				Response: zenzivaHomecareToPatientDone,
			}, {
				Msisdn:   superadmin.Phone,
				Action:   valHomecareToHealthOffice,
				Response: zenzivaHomecareToHealthOffice,
			}},
		)
	}

	// insert to transaction
	database.Datasource.DB().Create(
		&[]model.Transaction{{
			UserID:       hc.Chat.UserID,
			ChatID:       hc.ChatID,
			SystemStatus: statusHomecareToPatientDone.ValueSystem,
			NotifStatus:  notifMessageHomecareToPatientDone,
			UserStatus:   userMessageHomecareToPatientDone,
			PushStatus:   pushMessageHomecareToPatientDone,
		}, {
			UserID:       hc.Chat.UserID,
			ChatID:       hc.ChatID,
			SystemStatus: statusHomecareToHealthOffice.ValueSystem,
			NotifStatus:  notifMessageHomecareToHealthOffice,
			UserStatus:   userMessageHomecareToHealthOffice,
			PushStatus:   pushMessageHomecareToHealthOffice,
		}},
	)

	type resPushMessage struct {
		Patient      string `json:"patient"`
		HealthOffice string `json:"healthoffice"`
	}

	if req.RequestType == "mobile" {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":   false,
			"message": "Submited",
			"data":    homecareOfficer,
			"push_message": &resPushMessage{
				Patient:      pushMessageHomecareToPatientDone,
				HealthOffice: pushMessageHomecareToHealthOffice,
			},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
		"data":    homecareOfficer,
	})
}

func HomecarePhoto(c *fiber.Ctx) error {
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
			HomecareID: homecareId,
			Filename:   filename + "." + fileExt,
		},
	)

	// err := database.NewRedisClient().RPush().Err();
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Image uploaded successfully",
		"data":    imageFile,
	})
}
