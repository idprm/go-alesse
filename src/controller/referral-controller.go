package controller

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/config"
	"github.com/idprm/go-alesse/src/pkg/handler"
	"github.com/idprm/go-alesse/src/pkg/model"
)

type ReferralRequest struct {
	ChannelUrl   string `query:"channel_url" json:"channel_url"`
	SpecialistID int    `query:"specialist_id" json:"specialist_id"`
}

type ReferralChatRequest struct {
	ChannelUrl string `query:"channel_url" validate:"required" json:"channel_url"`
}

func Referral(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(ReferralRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var chat model.Chat
	database.Datasource.DB().Where("channel_url", req.ChannelUrl).Preload("Doctor").First(&chat)

	var specialist model.Specialist
	database.Datasource.DB().Where("id", req.SpecialistID).Where("is_active", true).First(&specialist)

	/**
	 * Check Referral
	 */
	var referral model.Referral
	resultReferral := database.Datasource.DB().
		Where("doctor_id", chat.Doctor.ID).
		Where("specialist_id", specialist.ID).
		Where("DATE(created_at) = DATE(?)", time.Now()).
		First(&referral)

	finishUrl := config.ViperEnv("APP_HOST") + "/specialist/chat"

	if resultReferral.RowsAffected == 0 {
		/**
		 * INSERT TO Referral
		 */
		database.Datasource.DB().Create(&model.Referral{
			SpecialistID: specialist.ID,
			DoctorID:     chat.DoctorID,
		})

		/**
		 * SENDBIRD PROCESS
		 */
		err := sendbirdProcessReferral(specialist.ID, chat.DoctorID)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
				"status":  fiber.StatusBadGateway,
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":        false,
			"message":      "Created Successfull",
			"redirect_url": finishUrl,
			"status":       fiber.StatusCreated,
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":        false,
			"message":      "Already chat",
			"redirect_url": finishUrl,
			"status":       fiber.StatusOK,
		})
	}
}

func ReferralChat(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(ReferralChatRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var referral model.Referral
	isReferral := database.Datasource.DB().Where("channel_url", referral.ChannelUrl).Preload("Specialist").Preload("Doctor").First(&referral)

	if isReferral.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&isReferral)
}

func sendbirdProcessReferral(specialistId uint, doctorId uint) error {

	var referral model.Referral
	database.Datasource.DB().
		Where("specialist_id", specialistId).
		Where("doctor_id", doctorId).
		Where("DATE(created_at) = DATE(?)", time.Now()).
		Preload("Doctor").
		Preload("Specialist").
		First(&referral)

		/**
		 * Check User Sendbird
		 */
	getSpecialist, _, err := handler.SendbirdGetSpecialist(referral.Specialist)
	if err != nil {
		return errors.New(err.Error())
	}

	/**
	 * Add User Sendbird
	 */
	database.Datasource.DB().Create(&model.Sendbird{
		Msisdn:   referral.Specialist.Phone,
		Action:   actionCheckUser,
		Response: getSpecialist,
	})

	// create group
	createGroup, name, url, err := handler.SendbirdReferralCreateGroupChannel(referral.Specialist, referral.Doctor)
	if err != nil {
		return errors.New(err.Error())
	}
	// insert to sendbird
	database.Datasource.DB().Create(&model.Sendbird{
		Msisdn:   referral.Doctor.Phone,
		Action:   actionCreateGroup,
		Response: createGroup,
	})

	if name != "" && url != "" {
		// insert to chat
		database.Datasource.DB().Create(
			&model.Referral{
				SpecialistID: referral.SpecialistID,
				DoctorID:     referral.DoctorID,
				ChannelName:  name,
				ChannelUrl:   url,
			})

		var conf model.Config
		database.Datasource.DB().Where("name", "NOTIF_MESSAGE_DOCTOR").First(&conf)

		urlWeb := config.ViperEnv("APP_HOST") + "/referral/" + url
		replaceMessage := strings.NewReplacer("@v1", referral.Specialist.Name, "@v2", referral.Doctor.Name, "@v3", urlWeb)
		message := replaceMessage.Replace(conf.Value)

		// NOTIF MESSAGE TO DOCTOR
		zenzifaNotif, err := handler.ZenzivaSendSMS(referral.Specialist.Phone, message)
		if err != nil {
			return errors.New(err.Error())
		}
		// insert to zenziva
		database.Datasource.DB().Create(&model.Zenziva{
			Msisdn:   referral.Specialist.Phone,
			Action:   actionCreateNotif,
			Response: zenzifaNotif,
		})
	}

	return nil
}
