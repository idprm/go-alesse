package controller

import (
	"log"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/config"
	"github.com/idprm/go-alesse/src/pkg/handler"
	"github.com/idprm/go-alesse/src/pkg/middleware"
	"github.com/idprm/go-alesse/src/pkg/model"
	"github.com/idprm/go-alesse/src/pkg/util"
)

type LoginRequest struct {
	Msisdn    string `query:"msisdn" validate:"required" json:"msisdn"`
	IpAddress string `query:"ip_address" json:"ip_address"`
}

type RegisterRequest struct {
	HealthCenter uint64 `query:"healthcenter_id" validate:"required" json:"healthcenter_id"`
	Msisdn       string `query:"msisdn" validate:"required" json:"msisdn"`
	Name         string `query:"name" validate:"required" json:"name"`
	Number       string `query:"number" validate:"required" json:"number"`
	Dob          string `query:"dob" validate:"required" json:"dob"`
	Gender       string `query:"gender" validate:"required" json:"gender"`
	Address      string `query:"address" validate:"required" json:"address"`
	Latitude     string `query:"latitude" json:"latitude"`
	Longitude    string `query:"longitude" json:"longitude"`
	UrlGmap      string `query:"url_gmap" json:"url_gmap"`
	IpAddress    string `query:"ip_address" json:"ip_address"`
}

type VerifyRequest struct {
	Msisdn    string `query:"msisdn" validate:"required" json:"msisdn"`
	Otp       string `query:"otp" validate:"required" json:"otp"`
	IpAddress string `query:"ip_address" json:"ip_address"`
}

type ErrorResponse struct {
	Field string
	Tag   string
	Value string
}

var validate = validator.New()

const (
	valOTPToUser  = "OTP_TO_USER"
	valOTPToAdmin = "OTP_TO_ADMIN"
)

func ValidateLogin(req LoginRequest) []*ErrorResponse {
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

func ValidateRegister(req RegisterRequest) []*ErrorResponse {
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

func ValidateVerify(req VerifyRequest) []*ErrorResponse {
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

func FrontHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":    false,
		"messsage": "Welcome to Alesse",
	})
}

func LoginHandler(c *fiber.Ctx) error {

	c.Accepts("application/json")

	req := new(LoginRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateLogin(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var user model.User
	isExist := database.Datasource.DB().Where("msisdn", req.Msisdn).Find(&user)

	if isExist.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Silakan daftar/registrasi",
		})
	}

	var usr model.User
	database.Datasource.DB().Where("msisdn", req.Msisdn).Preload("Healthcenter").First(&usr)
	usr.IpAddress = req.IpAddress
	usr.LoginAt = time.Now()
	database.Datasource.DB().Save(&usr)

	token, exp, err := middleware.GenerateJWTToken(usr)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	otp, _ := util.GenerateOTP(4)

	// insert to verify
	database.Datasource.DB().Create(
		&model.Verify{
			Msisdn:   req.Msisdn,
			Otp:      otp,
			IsVerify: false,
		},
	)

	var status model.Status
	database.Datasource.DB().Where("name", valOTPToUser).First(&status)
	notifMessageOTPToUser := util.ContentOTPToUser(status.ValueNotif, otp, config.ViperEnv("APP_HOST"))

	log.Println(notifMessageOTPToUser)

	zenzivaOTP, err := handler.ZenzivaSendSMS(req.Msisdn, notifMessageOTPToUser)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// insert to zenziva
	database.Datasource.DB().Create(&model.Zenziva{
		Msisdn:   user.Msisdn,
		Action:   valOTPToUser,
		Response: zenzivaOTP,
	})

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"error": false,
			"token": token,
			"exp":   exp,
			"user":  usr,
		},
	)
}

func RegisterHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(RegisterRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateRegister(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var user model.User
	isExist := database.Datasource.DB().Where("msisdn", req.Msisdn).First(&user)

	dob, _ := time.Parse("02-01-2006", req.Dob)

	if isExist.RowsAffected == 0 {
		database.Datasource.DB().Create(
			&model.User{
				HealthcenterID: req.HealthCenter,
				Msisdn:         req.Msisdn,
				Number:         req.Number,
				Name:           req.Name,
				Dob:            dob,
				Gender:         req.Gender,
				Address:        req.Address,
				Latitude:       req.Latitude,
				Longitude:      req.Longitude,
				UrlGmap:        req.UrlGmap,
				ActiveAt:       time.Now(),
				IsActive:       true,
			},
		)

		var usr model.User
		database.Datasource.DB().Where("msisdn", req.Msisdn).First(&usr)
		usr.IpAddress = req.IpAddress
		usr.LoginAt = time.Now()
		database.Datasource.DB().Save(&usr)

		token, exp, err := middleware.GenerateJWTToken(usr)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		otp, _ := util.GenerateOTP(4)

		// insert otp
		database.Datasource.DB().Create(
			&model.Verify{
				Msisdn:   req.Msisdn,
				Otp:      otp,
				IsVerify: false,
			},
		)

		var status model.Status
		database.Datasource.DB().Where("name", valOTPToUser).First(&status)
		notifMessageOTPToUser := util.ContentOTPToUser(status.ValueNotif, otp, config.ViperEnv("APP_HOST"))

		log.Println(notifMessageOTPToUser)

		zenzivaOTP, err := handler.ZenzivaSendSMS(req.Msisdn, notifMessageOTPToUser)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		// insert to zenziva
		database.Datasource.DB().Create(&model.Zenziva{
			Msisdn:   user.Msisdn,
			Action:   valOTPToUser,
			Response: zenzivaOTP,
		})

		return c.Status(fiber.StatusCreated).JSON(
			fiber.Map{
				"error": false,
				"token": token,
				"exp":   exp,
				"user":  usr,
			},
		)
	} else {
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"error":    false,
				"message":  "Already Active",
				"redirect": "/auth/login",
				"user":     user,
			},
		)
	}
}

func VerifyHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(VerifyRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateVerify(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var verify model.Verify
	checkOTP := database.Datasource.DB().Where("msisdn", req.Msisdn).Where("otp", req.Otp).Where("is_verify", false).First(&verify)

	if checkOTP.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Not found"})
	}

	if checkOTP.RowsAffected == 1 {
		// update status = 1 on verify
		database.Datasource.DB().Model(&verify).Where("msisdn", req.Msisdn).Where("otp", req.Otp).Update("is_verify", true)

		var user model.User
		database.Datasource.DB().Where("msisdn", req.Msisdn).First(&user)
		user.VerifyAt = time.Now()
		user.LoginAt = time.Now()
		user.IpAddress = req.IpAddress
		user.IsVerify = true
		database.Datasource.DB().Save(&user)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "data": verify})
}

func MLoginHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(LoginRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateLogin(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var user model.User
	isExist := database.Datasource.DB().Where("msisdn", req.Msisdn).Find(&user)

	if isExist.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Silakan daftar/registrasi",
		})
	}

	var usr model.User
	database.Datasource.DB().Where("msisdn", req.Msisdn).Preload("Healthcenter").First(&usr)
	usr.IpAddress = req.IpAddress
	usr.LoginAt = time.Now()
	database.Datasource.DB().Save(&usr)

	token, exp, err := middleware.GenerateJWTToken(usr)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	otp, _ := util.GenerateOTP(4)

	// insert to verify
	database.Datasource.DB().Create(
		&model.Verify{
			Msisdn:   req.Msisdn,
			Otp:      otp,
			IsVerify: false,
		},
	)

	var status model.Status
	database.Datasource.DB().Where("name", valOTPToUser).First(&status)
	notifMessageOTPToUser := util.ContentOTPToUser(status.ValueNotif, otp, config.ViperEnv("APP_HOST"))

	log.Println(notifMessageOTPToUser)

	zenzivaOTP, err := handler.ZenzivaSendSMS(req.Msisdn, notifMessageOTPToUser)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// insert to zenziva
	database.Datasource.DB().Create(&model.Zenziva{
		Msisdn:   user.Msisdn,
		Action:   valOTPToUser,
		Response: zenzivaOTP,
	})

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"error": false,
			"token": token,
			"exp":   exp,
			"user":  usr,
		},
	)
}

func MRegisterHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(RegisterRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateRegister(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var user model.User
	isExist := database.Datasource.DB().Where("msisdn", req.Msisdn).First(&user)

	dob, _ := time.Parse("02-01-2006", req.Dob)

	if isExist.RowsAffected == 0 {
		database.Datasource.DB().Create(
			&model.User{
				HealthcenterID: req.HealthCenter,
				Msisdn:         req.Msisdn,
				Number:         req.Number,
				Name:           req.Name,
				Dob:            dob,
				Gender:         req.Gender,
				Address:        req.Address,
				Latitude:       req.Latitude,
				Longitude:      req.Longitude,
				ActiveAt:       time.Now(),
				IsActive:       true,
			},
		)

		var usr model.User
		database.Datasource.DB().Where("msisdn", req.Msisdn).First(&usr)
		usr.IpAddress = req.IpAddress
		usr.LoginAt = time.Now()
		database.Datasource.DB().Save(&usr)

		token, exp, err := middleware.GenerateJWTToken(usr)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		otp, _ := util.GenerateOTP(4)

		// insert otp
		database.Datasource.DB().Create(
			&model.Verify{
				Msisdn:   req.Msisdn,
				Otp:      otp,
				IsVerify: false,
			},
		)

		var status model.Status
		database.Datasource.DB().Where("name", valOTPToUser).First(&status)
		notifMessageOTPToUser := util.ContentOTPToUser(status.ValueNotif, otp, config.ViperEnv("APP_HOST"))

		log.Println(notifMessageOTPToUser)

		zenzivaOTP, err := handler.ZenzivaSendSMS(req.Msisdn, notifMessageOTPToUser)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		// insert to zenziva
		database.Datasource.DB().Create(&model.Zenziva{
			Msisdn:   user.Msisdn,
			Action:   valOTPToUser,
			Response: zenzivaOTP,
		})

		return c.Status(fiber.StatusCreated).JSON(
			fiber.Map{
				"error": false,
				"token": token,
				"exp":   exp,
				"user":  usr,
			},
		)
	} else {
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"error":    false,
				"message":  "Already Active",
				"redirect": "/auth/login",
				"user":     user,
			},
		)
	}
}

func MVerifyHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(VerifyRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateVerify(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var verify model.Verify
	checkOTP := database.Datasource.DB().Where("msisdn", req.Msisdn).Where("otp", req.Otp).Where("is_verify", false).First(&verify)

	if checkOTP.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Not found"})
	}

	if checkOTP.RowsAffected == 1 {
		// update status = 1 on verify
		database.Datasource.DB().Model(&verify).Where("msisdn", req.Msisdn).Where("otp", req.Otp).Update("is_verify", true)

		var user model.User
		database.Datasource.DB().Where("msisdn", req.Msisdn).First(&user)
		user.VerifyAt = time.Now()
		user.LoginAt = time.Now()
		user.IpAddress = req.IpAddress
		user.IsVerify = true
		database.Datasource.DB().Save(&user)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "data": verify})
}

func MAuthDoctorHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(LoginRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateLogin(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var doctor model.Doctor
	isExist := database.Datasource.DB().Where("phone", req.Msisdn).Find(&doctor)

	if isExist.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Not found",
		})
	}

	var dtr model.Doctor
	database.Datasource.DB().Where("phone", req.Msisdn).Preload("Healthcenter").First(&dtr)
	dtr.IpAddress = req.IpAddress
	dtr.LoginAt = time.Now()
	database.Datasource.DB().Save(&dtr)
	token, exp, err := middleware.GenerateJWTTokenDoctor(dtr)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	otp, _ := util.GenerateOTP(4)

	// insert to verify
	database.Datasource.DB().Create(
		&model.Verify{
			Msisdn:   req.Msisdn,
			Otp:      otp,
			IsVerify: false,
		},
	)

	var status model.Status
	database.Datasource.DB().Where("name", valOTPToAdmin).First(&status)
	notifMessageOTPToUser := util.ContentOTPToUser(status.ValueNotif, otp, config.ViperEnv("APP_HOST"))

	zenzivaOTP, err := handler.ZenzivaSendSMS(req.Msisdn, notifMessageOTPToUser)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// insert to zenziva
	database.Datasource.DB().Create(&model.Zenziva{
		Msisdn:   doctor.Phone,
		Action:   valOTPToUser,
		Response: zenzivaOTP,
	})

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"error": false,
			"token": token,
			"exp":   exp,
			"data":  dtr,
		},
	)
}

func MAuthOfficerHandler(c *fiber.Ctx) error {
	req := new(LoginRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateLogin(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var officer model.Officer
	isExist := database.Datasource.DB().Where("phone", req.Msisdn).Find(&officer)

	if isExist.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Not found",
		})
	}

	var ofcr model.Officer
	database.Datasource.DB().Where("phone", req.Msisdn).Preload("Healthcenter").First(&ofcr)
	ofcr.IpAddress = req.IpAddress
	ofcr.LoginAt = time.Now()
	database.Datasource.DB().Save(&ofcr)
	token, exp, err := middleware.GenerateJWTTokenOfficer(ofcr)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	otp, _ := util.GenerateOTP(4)

	// insert to verify
	database.Datasource.DB().Create(
		&model.Verify{
			Msisdn:   req.Msisdn,
			Otp:      otp,
			IsVerify: false,
		},
	)

	var status model.Status
	database.Datasource.DB().Where("name", valOTPToAdmin).First(&status)
	notifMessageOTPToUser := util.ContentOTPToUser(status.ValueNotif, otp, config.ViperEnv("APP_HOST"))

	zenzivaOTP, err := handler.ZenzivaSendSMS(req.Msisdn, notifMessageOTPToUser)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// insert to zenziva
	database.Datasource.DB().Create(&model.Zenziva{
		Msisdn:   ofcr.Phone,
		Action:   valOTPToUser,
		Response: zenzivaOTP,
	})

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"error": false,
			"token": token,
			"exp":   exp,
			"data":  ofcr,
		},
	)
}

func MAuthApothecaryHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(LoginRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateLogin(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var apothecary model.Apothecary
	isExist := database.Datasource.DB().Where("phone", req.Msisdn).Find(&apothecary)

	if isExist.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Not found",
		})
	}

	var apot model.Apothecary
	database.Datasource.DB().Where("phone", req.Msisdn).Preload("Healthcenter").First(&apot)
	apot.IpAddress = req.IpAddress
	apot.LoginAt = time.Now()
	database.Datasource.DB().Save(&apot)
	token, exp, err := middleware.GenerateJWTTokenApothecary(apot)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	otp, _ := util.GenerateOTP(4)

	// insert to verify
	database.Datasource.DB().Create(
		&model.Verify{
			Msisdn:   req.Msisdn,
			Otp:      otp,
			IsVerify: false,
		},
	)

	var status model.Status
	database.Datasource.DB().Where("name", valOTPToAdmin).First(&status)
	notifMessageOTPToUser := util.ContentOTPToUser(status.ValueNotif, otp, config.ViperEnv("APP_HOST"))

	zenzivaOTP, err := handler.ZenzivaSendSMS(req.Msisdn, notifMessageOTPToUser)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// insert to zenziva
	database.Datasource.DB().Create(&model.Zenziva{
		Msisdn:   apot.Phone,
		Action:   valOTPToUser,
		Response: zenzivaOTP,
	})

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"error": false,
			"token": token,
			"exp":   exp,
			"data":  apot,
		},
	)
}

func MAuthCourierHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(LoginRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateLogin(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var courier model.Courier
	isExist := database.Datasource.DB().Where("phone", req.Msisdn).Find(&courier)

	if isExist.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Not found",
		})
	}

	var cour model.Courier
	database.Datasource.DB().Where("phone", req.Msisdn).Preload("Healthcenter").First(&cour)
	cour.IpAddress = req.IpAddress
	cour.LoginAt = time.Now()
	database.Datasource.DB().Save(&cour)
	token, exp, err := middleware.GenerateJWTTokenCourier(cour)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	otp, _ := util.GenerateOTP(4)

	// insert to verify
	database.Datasource.DB().Create(
		&model.Verify{
			Msisdn:   req.Msisdn,
			Otp:      otp,
			IsVerify: false,
		},
	)

	var status model.Status
	database.Datasource.DB().Where("name", valOTPToAdmin).First(&status)
	notifMessageOTPToUser := util.ContentOTPToUser(status.ValueNotif, otp, config.ViperEnv("APP_HOST"))

	zenzivaOTP, err := handler.ZenzivaSendSMS(req.Msisdn, notifMessageOTPToUser)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// insert to zenziva
	database.Datasource.DB().Create(&model.Zenziva{
		Msisdn:   cour.Phone,
		Action:   valOTPToUser,
		Response: zenzivaOTP,
	})

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"error": false,
			"token": token,
			"exp":   exp,
			"data":  cour,
		},
	)
}

func MAuthSpecialistHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(LoginRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateLogin(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var specialist model.Specialist
	isExist := database.Datasource.DB().Where("phone", req.Msisdn).Find(&specialist)

	if isExist.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Not found",
		})
	}

	var spec model.Specialist
	database.Datasource.DB().Where("phone", req.Msisdn).First(&spec)
	spec.IpAddress = req.IpAddress
	spec.LoginAt = time.Now()
	database.Datasource.DB().Save(&spec)
	token, exp, err := middleware.GenerateJWTTokenSpecialist(spec)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	otp, _ := util.GenerateOTP(4)

	// insert to verify
	database.Datasource.DB().Create(
		&model.Verify{
			Msisdn:   req.Msisdn,
			Otp:      otp,
			IsVerify: false,
		},
	)

	var status model.Status
	database.Datasource.DB().Where("name", valOTPToAdmin).First(&status)
	notifMessageOTPToUser := util.ContentOTPToUser(status.ValueNotif, otp, config.ViperEnv("APP_HOST"))

	zenzivaOTP, err := handler.ZenzivaSendSMS(req.Msisdn, notifMessageOTPToUser)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// insert to zenziva
	database.Datasource.DB().Create(&model.Zenziva{
		Msisdn:   spec.Phone,
		Action:   valOTPToUser,
		Response: zenzivaOTP,
	})

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"error": false,
			"token": token,
			"exp":   exp,
			"data":  spec,
		},
	)
}
