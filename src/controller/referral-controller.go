package controller

import (
	"errors"
	"strings"

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
	resultReferral := database.Datasource.DB().Where("chat_id", chat.ID).First(&referral)

	finishUrl := config.ViperEnv("APP_HOST") + "/referral/" + referral.ChannelUrl

	if resultReferral.RowsAffected == 0 {

		/**
		 * SENDBIRD PROCESS
		 */
		channelUrl, err := sendbirdProcessReferral(chat.ID, specialist.ID, chat.DoctorID)
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
			"redirect_url": config.ViperEnv("APP_HOST") + "/referral/" + channelUrl,
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

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var referral model.Referral
	database.Datasource.DB().Where("channel_url", req.ChannelUrl).Preload("Specialist").Preload("Doctor").First(&referral)

	return c.Status(fiber.StatusOK).JSON(referral)
}

func sendbirdProcessReferral(chatId uint64, specialistId uint, doctorId uint) (string, error) {

	var specialist model.Specialist
	database.Datasource.DB().Where("id", specialistId).First(&specialist)

	var doctor model.Doctor
	database.Datasource.DB().Where("id", doctorId).First(&doctor)

	/**
	 * Check User Sendbird
	 */
	getSpecialist, _, err := handler.SendbirdGetSpecialist(specialist)
	if err != nil {
		return "", errors.New(err.Error())
	}

	/**
	 * Add User Sendbird
	 */
	database.Datasource.DB().Create(&model.Sendbird{
		Msisdn:   specialist.Phone,
		Action:   actionCheckUser,
		Response: getSpecialist,
	})

	// create group
	createGroup, name, url, err := handler.SendbirdReferralCreateGroupChannel(specialist, doctor)
	if err != nil {
		return "", errors.New(err.Error())
	}
	// insert to sendbird
	database.Datasource.DB().Create(&model.Sendbird{
		Msisdn:   doctor.Phone,
		Action:   actionCreateGroup,
		Response: createGroup,
	})

	if name != "" && url != "" {
		// insert to chat referral
		database.Datasource.DB().Create(
			&model.Referral{
				ChatID:       chatId,
				SpecialistID: specialistId,
				DoctorID:     doctorId,
				ChannelName:  name,
				ChannelUrl:   url,
			})

		var conf model.Config
		database.Datasource.DB().Where("name", "NOTIF_MESSAGE_SPECIALIST").First(&conf)

		urlWeb := config.ViperEnv("APP_HOST") + "/specialist/chat/" + url
		replaceMessage := strings.NewReplacer("@v1", specialist.Name, "@v2", doctor.Name, "@v3", urlWeb)
		message := replaceMessage.Replace(conf.Value)

		// NOTIF MESSAGE TO SPECIALIST
		zenzifaNotif, err := handler.ZenzivaSendSMS(specialist.Phone, message)
		if err != nil {
			return "", errors.New(err.Error())
		}
		// insert to zenziva
		database.Datasource.DB().Create(&model.Zenziva{
			Msisdn:   specialist.Phone,
			Action:   actionCreateNotifSP,
			Response: zenzifaNotif,
		})

	}

	return url, nil
}
