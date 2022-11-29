package controller

import (
	"errors"
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

type PharmacyRequest struct {
	ChatID          uint64                    `query:"chat_id" validate:"required" json:"chat_id"`
	Weight          uint32                    `query:"weight" validate:"required" json:"weight"`
	PainComplaints  string                    `query:"pain_complaints" validate:"required" json:"pain_complaints"`
	Diagnosis       string                    `query:"diagnosis" validate:"required" json:"diagnosis"`
	AllergyMedicine string                    `query:"allergy_medicine" validate:"required" json:"allergy_medicine"`
	Slug            string                    `query:"slug" json:"slug"`
	Data            []PharmacyMedicineRequest `query:"data" json:"data"`
}

type PharmacyMedicineRequest struct {
	PharmacyID  uint64 `query:"pharmacy_id" json:"pharmacy_id"`
	MedicineID  uint64 `query:"medicine_id" json:"medicine_id"`
	Name        string `query:"name" json:"name"`
	Qty         uint   `query:"quantity" json:"quantity"`
	Rule        string `query:"rule" json:"rule"`
	Dose        string `query:"dose" json:"dose"`
	Description string `query:"description" json:"description"`
}

type PharmacyProcessRequest struct {
	PharmacyID uint64 `query:"pharmacy_id" validate:"required" json:"pharmacy_id"`
}

type PharmacySentRequest struct {
	PharmacyID uint64 `query:"pharmacy_id" validate:"required" json:"pharmacy_id"`
}

type PharmacyTakeRequest struct {
	PharmacyID uint64 `query:"pharmacy_id" validate:"required" json:"pharmacy_id"`
}

type PharmacyFinishRequest struct {
	PharmacyID uint64 `query:"pharmacy_id" validate:"required" json:"pharmacy_id"`
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

func ValidatePharmacyTake(req PharmacyTakeRequest) []*ErrorResponse {
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

func ValidatePharmacyFinish(req PharmacyFinishRequest) []*ErrorResponse {
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

func GetPharmacy(c *fiber.Ctx) error {
	var pharmacy model.Pharmacy
	database.Datasource.DB().Where("slug", c.Params("slug")).First(&pharmacy)
	return c.Status(fiber.StatusOK).JSON(pharmacy)
}

func GetPharmacyByDoctor(c *fiber.Ctx) error {
	var pharmacy model.Pharmacy
	database.Datasource.DB().Where("slug", c.Params("slug")).Preload("Chat.Doctor").Preload("Chat.User").First(&pharmacy)
	return c.Status(fiber.StatusOK).JSON(pharmacy)
}

func GetAllPharmacyByMedicines(c *fiber.Ctx) error {
	var pharmacy model.Pharmacy
	database.Datasource.DB().Where("slug", c.Params("slug")).First(&pharmacy)

	var medicines []model.PharmacyMedicine
	database.Datasource.DB().Where("pharmacy_id", pharmacy.ID).Preload("Medicine").Find(&medicines)
	return c.Status(fiber.StatusOK).JSON(medicines)
}

func GetPharmacyByApothecary(c *fiber.Ctx) error {
	var pharmacy model.Pharmacy
	database.Datasource.DB().Where("slug", c.Params("slug")).First(&pharmacy)

	var pharmacyApothecary model.PharmacyApothecary
	database.Datasource.DB().Where("pharmacy_id", pharmacy.ID).Preload("Pharmacy.Chat.Doctor").Preload("Pharmacy.Chat.User").First(&pharmacyApothecary)
	return c.Status(fiber.StatusOK).JSON(pharmacyApothecary)
}

func GetPharmacyByCourier(c *fiber.Ctx) error {
	var pharmacy model.Pharmacy
	database.Datasource.DB().Where("slug", c.Params("slug")).First(&pharmacy)

	var pharmacyCourier model.PharmacyCourier
	database.Datasource.DB().Where("pharmacy_id", pharmacy.ID).Preload("Pharmacy.Chat.Doctor").Preload("Pharmacy.Chat.User").First(&pharmacyCourier)
	return c.Status(fiber.StatusOK).JSON(pharmacyCourier)
}

func GetPharmacyAllPhoto(c *fiber.Ctx) error {
	var photos []model.Photo
	database.Datasource.DB().Where("pharmacy_id", c.Query("id")).Find(&photos)
	return c.Status(fiber.StatusOK).JSON(photos)
}

func SavePharmacy(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(PharmacyRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var pharmacy model.Pharmacy
	isExist := database.Datasource.DB().Where("chat_id", req.ChatID).First(&pharmacy)

	if isExist.RowsAffected == 0 {
		pharmacy := model.Pharmacy{
			ChatID:          req.ChatID,
			Weight:          req.Weight,
			PainComplaints:  req.PainComplaints,
			Diagnosis:       req.Diagnosis,
			AllergyMedicine: req.AllergyMedicine,
			Slug:            req.Slug,
			Number:          util.TimeStamp(),
			SubmitedAt:      time.Now(),
			IsSubmited:      true,
		}
		database.Datasource.DB().Create(&pharmacy)

		database.Datasource.DB().Where("pharmacy_id", pharmacy.ID).Delete(&model.PharmacyMedicine{})

		for _, v := range req.Data {
			database.Datasource.DB().Create(&model.PharmacyMedicine{
				PharmacyID:  pharmacy.ID,
				MedicineID:  v.MedicineID,
				Name:        v.Name,
				Quantity:    v.Qty,
				Rule:        v.Rule,
				Dose:        v.Dose,
				Description: v.Description,
			})
		}

	} else {
		pharmacy.Weight = req.Weight
		pharmacy.PainComplaints = req.PainComplaints
		pharmacy.Diagnosis = req.Diagnosis
		pharmacy.AllergyMedicine = req.AllergyMedicine
		pharmacy.Slug = req.Slug
		pharmacy.SubmitedAt = time.Now()
		pharmacy.IsSubmited = true
		database.Datasource.DB().Save(&pharmacy)

		database.Datasource.DB().Where("pharmacy_id", pharmacy.ID).Delete(&model.PharmacyMedicine{})

		for _, v := range req.Data {
			database.Datasource.DB().Create(&model.PharmacyMedicine{
				PharmacyID:  pharmacy.ID,
				MedicineID:  v.MedicineID,
				Name:        v.Name,
				Quantity:    v.Qty,
				Rule:        v.Rule,
				Dose:        v.Dose,
				Description: v.Description,
			})
		}

	}

	/**
	 * NOTIF_DOCTOR_TO_PHARMACY
	 */
	const (
		valDoctorToPharmacy  = "NOTIF_DOCTOR_TO_PHARMACY"
		valPharmacyToPatient = "NOTIF_PHARMACY_TO_PATIENT"
	)

	var phar model.Pharmacy
	database.Datasource.DB().Where("chat_id", req.ChatID).Preload("Chat.User").Preload("Chat.Doctor").Preload("Chat.Healthcenter").First(&phar)

	var apothecary model.Apothecary
	database.Datasource.DB().Where("healthcenter_id", phar.Chat.HealthcenterID).First(&apothecary)

	var confDoctorToPharmacy model.Config
	database.Datasource.DB().Where("name", valDoctorToPharmacy).First(&confDoctorToPharmacy)
	replaceMessageDoctorToPharmacy := util.ContentDoctorToPharmacy(confDoctorToPharmacy.Value, phar)

	log.Println(replaceMessageDoctorToPharmacy)

	var confPharmacyToPatient model.Config
	database.Datasource.DB().Where("name", valPharmacyToPatient).First(&confPharmacyToPatient)
	replaceMessagePharmacyToPatient := util.ContentDoctorToPharmacy(confPharmacyToPatient.Value, phar)

	log.Println(replaceMessagePharmacyToPatient)

	zenzifaNotifDoctorToPharmacy, err := handler.ZenzivaSendSMS(apothecary.Phone, replaceMessageDoctorToPharmacy)
	if err != nil {
		return errors.New(err.Error())
	}

	zenzifaNotifPharmacyToPatient, err := handler.ZenzivaSendSMS(phar.Chat.User.Msisdn, replaceMessagePharmacyToPatient)
	if err != nil {
		return errors.New(err.Error())
	}

	// insert to zenziva
	database.Datasource.DB().Create(&model.Zenziva{
		Msisdn:   apothecary.Phone,
		Action:   valDoctorToPharmacy,
		Response: zenzifaNotifDoctorToPharmacy,
	})

	// insert to zenziva
	database.Datasource.DB().Create(&model.Zenziva{
		Msisdn:   phar.Chat.User.Msisdn,
		Action:   valPharmacyToPatient,
		Response: zenzifaNotifPharmacyToPatient,
	})

	// insert to notif
	database.Datasource.DB().Create(
		&model.Notif{
			UserID:  phar.Chat.UserID,
			Content: "",
		},
	)

	// insert to notif
	database.Datasource.DB().Create(
		&model.Notif{
			UserID:  phar.Chat.UserID,
			Content: "",
		},
	)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
		"data":    pharmacy,
	})
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

	var pharmacyApothecary model.PharmacyApothecary
	existPharmacyApothecary := database.Datasource.DB().Where("pharmacy_id", req.PharmacyID).First(&pharmacyApothecary)

	if existPharmacyApothecary.RowsAffected == 0 {
		pharmacyApothecary := model.PharmacyApothecary{
			PharmacyID: req.PharmacyID,
			ProcessAt:  time.Now(),
			IsProcess:  true,
		}
		database.Datasource.DB().Create(&pharmacyApothecary)
	} else {
		pharmacyApothecary.ProcessAt = time.Now()
		pharmacyApothecary.IsProcess = true
		database.Datasource.DB().Save(&pharmacyApothecary)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
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

	var pharmacyApothecary model.PharmacyApothecary
	existPharmacy := database.Datasource.DB().Where("pharmacy_id", req.PharmacyID).First(&pharmacyApothecary)

	if existPharmacy.RowsAffected > 0 {

		pharmacyApothecary.SentAt = time.Now()
		pharmacyApothecary.IsSent = true
		database.Datasource.DB().Save(&pharmacyApothecary)

		// NOTIF_PHARMACY_TO_COURIER
		var pharmacy model.Pharmacy
		database.Datasource.DB().Where("id", req.PharmacyID).Preload("Chat.User").Preload("Chat.Doctor").Preload("Chat.Healthcenter").First(&pharmacy)

		var courier model.Courier
		database.Datasource.DB().Where("healthcenter_id", pharmacy.Chat.HealthcenterID).First(&courier)

		const (
			valPharmacyToCourier = "NOTIF_PHARMACY_TO_COURIER"
		)
		var conf model.Config
		database.Datasource.DB().Where("name", valPharmacyToCourier).First(&conf)
		replaceMessage := util.ContentPharmacyToCourier(conf.Value, pharmacy, courier)

		zenzivaPharmacyToCourier, err := handler.ZenzivaSendSMS(courier.Phone, replaceMessage)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		// insert to zenziva
		database.Datasource.DB().Create(
			&model.Zenziva{
				Msisdn:   courier.Phone,
				Action:   valPharmacyToCourier,
				Response: zenzivaPharmacyToCourier,
			},
		)

		// insert to notif
		database.Datasource.DB().Create(
			&model.Notif{
				UserID:  pharmacy.Chat.User.ID,
				Content: "",
			},
		)

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
		"data":    pharmacyApothecary,
	})
}

func SaveTakePharmacy(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(PharmacyTakeRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidatePharmacyTake(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var pharmacyCourier model.PharmacyCourier
	existPharmacy := database.Datasource.DB().Where("pharmacy_id", req.PharmacyID).First(&pharmacyCourier)

	if existPharmacy.RowsAffected == 0 {
		pharmacyCourier := model.PharmacyCourier{
			PharmacyID: req.PharmacyID,
			TakeAt:     time.Now(),
			IsTake:     true,
		}
		database.Datasource.DB().Create(&pharmacyCourier)

		const (
			valCourierToPatient = "NOTIF_COURIER_TO_PATIENT"
		)

		// NOTIF_COURIER_TO_PHARMACY
		var pharmacy model.Pharmacy
		database.Datasource.DB().Where("id", req.PharmacyID).Preload("Chat.User").Preload("Chat.Doctor").Preload("Chat.Healthcenter").First(&pharmacy)

		var confCourierToPatient model.Config
		database.Datasource.DB().Where("name", valCourierToPatient).First(&confCourierToPatient)
		replaceMessageCourierToPatient := util.ContentCourierToPatient(confCourierToPatient.Value, pharmacy)

		zenzivaCourierToPatient, err := handler.ZenzivaSendSMS(pharmacy.Chat.User.Msisdn, replaceMessageCourierToPatient)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		// insert to zenziva
		database.Datasource.DB().Create(
			&model.Zenziva{
				Msisdn:   pharmacy.Chat.User.Msisdn,
				Action:   valCourierToPatient,
				Response: zenzivaCourierToPatient,
			},
		)

		// insert to notif
		database.Datasource.DB().Create(
			&model.Notif{
				UserID:  pharmacy.Chat.UserID,
				Content: "",
			},
		)

	} else {
		pharmacyCourier.TakeAt = time.Now()
		pharmacyCourier.IsTake = true
		database.Datasource.DB().Save(&pharmacyCourier)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
	})
}

func SaveFinishPharmacy(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(PharmacyFinishRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidatePharmacyFinish(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var pharmacyCourier model.PharmacyCourier
	existPharmacy := database.Datasource.DB().Where("pharmacy_id", req.PharmacyID).First(&pharmacyCourier)

	if existPharmacy.RowsAffected > 0 {
		pharmacyCourier.FinishAt = time.Now()
		pharmacyCourier.IsFinish = true
		database.Datasource.DB().Save(&pharmacyCourier)

		// NOTIF_COURIER_TO_PHARMACY
		var pharmacy model.Pharmacy
		database.Datasource.DB().Where("id", req.PharmacyID).Preload("Chat").Preload("Chat.User").Preload("Chat.Doctor").Preload("Chat.Healthcenter").First(&pharmacy)
		var courier model.Courier
		database.Datasource.DB().Where("healthcenter_id", pharmacy.Chat.HealthcenterID).First(&courier)

		const (
			valCourierToPharmacy = "NOTIF_COURIER_TO_PHARMACY"
			valFeedbackToPatient = "NOTIF_FEEDBACK_TO_PATIENT"
		)
		var confCourierToPharmacy model.Config
		database.Datasource.DB().Where("name", valCourierToPharmacy).First(&confCourierToPharmacy)
		replaceMessageCourierToPharmacy := util.ContentCourierToPharmacy(confCourierToPharmacy.Value, pharmacy, courier)
		log.Println(replaceMessageCourierToPharmacy)

		var confFeedbackToPatient model.Config
		database.Datasource.DB().Where("name", valFeedbackToPatient).First(&confFeedbackToPatient)
		replaceMessageFeedbackToPatient := util.ContentFeedbackToPatient(confFeedbackToPatient.Value, pharmacy.Chat)
		log.Println(replaceMessageFeedbackToPatient)

		zenzivaPharmacyToCourier, err := handler.ZenzivaSendSMS(courier.Phone, replaceMessageCourierToPharmacy)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		zenzivaFeedbackToPatient, err := handler.ZenzivaSendSMS(pharmacy.Chat.User.Msisdn, replaceMessageFeedbackToPatient)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		// insert to zenziva
		database.Datasource.DB().Create(
			&model.Zenziva{
				Msisdn:   courier.Phone,
				Action:   valCourierToPharmacy,
				Response: zenzivaPharmacyToCourier,
			},
		)

		// insert to zenziva
		database.Datasource.DB().Create(
			&model.Zenziva{
				Msisdn:   pharmacy.Chat.User.Msisdn,
				Action:   valFeedbackToPatient,
				Response: zenzivaFeedbackToPatient,
			},
		)

		// insert to notif
		database.Datasource.DB().Create(
			&model.Notif{
				UserID:  pharmacy.Chat.User.ID,
				Content: "",
			},
		)

		// insert to notif
		database.Datasource.DB().Create(
			&model.Notif{
				UserID:  pharmacy.Chat.User.ID,
				Content: "",
			},
		)

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
	})
}

func PharmacyPhoto(c *fiber.Ctx) error {
	pharmacyId, _ := strconv.ParseUint(c.FormValue("pharmacy_id"), 0, 64)
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
	err = c.SaveFile(file, fmt.Sprintf("./public/uploads/pharmacy/%s", imageFile))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   false,
			"message": "Server error",
			"data":    nil,
		})
	}

	database.Datasource.DB().Create(
		&model.PharmacyUpload{
			PharmacyID: pharmacyId,
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
