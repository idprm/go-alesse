package controller

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/middleware"
	"github.com/idprm/go-alesse/src/pkg/model"
)

type AuthRequest struct {
	Msisdn       string `query:"msisdn" validate:"required" json:"msisdn"`
	Name         string `query:"name" validate:"required" json:"name"`
	HealthCenter uint64 `query:"healthcenter_id" validate:"required" json:"healthcenter_id"`
}

type VerifyRequest struct {
	Msisdn string `query:"msisdn" validate:"required" json:"msisdn"`
	Otp    string `query:"otp" validate:"required" json:"otp"`
}

type ErrorResponse struct {
	Field string
	Tag   string
	Value string
}

var validate = validator.New()

func ValidateAuth(req AuthRequest) []*ErrorResponse {
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

func AuthHandler(c *fiber.Ctx) error {

	c.Accepts("application/json")

	req := new(AuthRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateAuth(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var user model.User
	isExist := database.Datasource.DB().Where("msisdn", req.Msisdn).First(&user)

	if isExist.RowsAffected == 0 {
		database.Datasource.DB().Create(&model.User{
			Msisdn:         req.Msisdn,
			Name:           req.Name,
			HealthcenterID: req.HealthCenter,
		})
	} else {
		user.Msisdn = req.Msisdn
		user.Name = req.Name
		user.HealthcenterID = req.HealthCenter
		database.Datasource.DB().Save(&user)
	}

	var usr model.User
	database.Datasource.DB().Where("msisdn", req.Msisdn).First(&usr)

	token, exp, err := middleware.GenerateJWTToken(usr)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	verify := model.Verify{Msisdn: req.Msisdn, Otp: "1234", IsVerify: false}
	database.Datasource.DB().Create(&verify)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"token": token,
		"exp":   exp,
		"user":  usr,
	})
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
	checkOTP := database.Datasource.DB().Where("msisdn", req.Msisdn).Where("otp", req.Otp).First(&verify)

	if checkOTP.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Not found"})
	}

	if checkOTP.RowsAffected == 1 {
		// update status = 1 on verify
		database.Datasource.DB().Model(&verify).Where("msisdn", req.Msisdn).Where("otp", req.Otp).Update("is_verify", true)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "data": verify})
}
