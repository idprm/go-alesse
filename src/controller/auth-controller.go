package controller

import (
	"log"
	"strings"
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

	// insert to verify
	database.Datasource.DB().Create(
		&model.Verify{
			Msisdn:   req.Msisdn,
			Otp:      otp,
			IsVerify: false,
		},
	)

	var confNotifOTP model.Config
	database.Datasource.DB().Where("name", "NOTIF_OTP_USER").First(&confNotifOTP)
	replaceMessageNotifOTP := strings.NewReplacer("@v1", otp)
	contentNotifOTP := replaceMessageNotifOTP.Replace(confNotifOTP.Value)

	log.Println(contentNotifOTP)

	zenzivaOTP, err := handler.ZenzivaSendSMS(req.Msisdn, contentNotifOTP)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// insert to zenziva
	database.Datasource.DB().Create(&model.Zenziva{
		Msisdn:   user.Msisdn,
		Action:   "OTP",
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

		const (
			valOTPToUser = "OTP_TO_USER"
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
