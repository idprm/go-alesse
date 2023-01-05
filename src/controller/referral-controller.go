package controller

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/config"
	"github.com/idprm/go-alesse/src/pkg/handler"
	"github.com/idprm/go-alesse/src/pkg/model"
	"github.com/idprm/go-alesse/src/pkg/util"
)

type ReferralRequest struct {
	RequestType  string `query:"request_type" validate:"required" json:"request_type"`
	ChannelUrl   string `query:"channel_url" json:"channel_url"`
	SpecialistID int    `query:"specialist_id" json:"specialist_id"`
}

type ReferralChatRequest struct {
	ChannelUrl string `query:"channel_url" validate:"required" json:"channel_url"`
}

const (
	valMessageToSpecialist = "MESSAGE_TO_SPECIALIST"
)

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
		channelUrl, pushMessage, err := sendbirdProcessReferral(chat.ID, specialist.ID, chat.DoctorID, req.RequestType)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
				"status":  fiber.StatusBadGateway,
			})
		}

		var status model.Status
		database.Datasource.DB().Where("name", valMessageToSpecialist).First(&status)

		if req.RequestType == "mobile" {
			return c.Status(fiber.StatusCreated).JSON(fiber.Map{
				"error":        false,
				"code":         fiber.StatusCreated,
				"message":      "Created Successfull",
				"channel_url":  channelUrl,
				"push_message": pushMessage,
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":        false,
			"message":      "Created Successfull",
			"redirect_url": config.ViperEnv("APP_HOST") + "/referral/" + channelUrl,
			"status":       fiber.StatusCreated,
		})

	} else {
		if req.RequestType == "mobile" {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"error":       false,
				"message":     "Already chat",
				"channel_url": referral.ChannelUrl,
				"status":      fiber.StatusOK,
			})
		}
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
	database.Datasource.DB().Where("channel_url", req.ChannelUrl).Preload("Chat").Preload("Specialist").Preload("Doctor").First(&referral)

	return c.Status(fiber.StatusOK).JSON(referral)
}

func sendbirdProcessReferral(chatId uint64, specialistId uint, doctorId uint, requestType string) (string, string, error) {

	var chat model.Chat
	database.Datasource.DB().Where("id", chatId).First(&chat)

	var specialist model.Specialist
	database.Datasource.DB().Where("id", specialistId).First(&specialist)

	var doctor model.Doctor
	database.Datasource.DB().Where("id", doctorId).First(&doctor)

	/**
	 * Check User Sendbird
	 */
	getSpecialist, _, err := handler.SendbirdGetSpecialist(specialist)
	if err != nil {
		return "", "", errors.New(err.Error())
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
		return "", "", errors.New(err.Error())
	}
	// insert to sendbird
	database.Datasource.DB().Create(&model.Sendbird{
		Msisdn:   doctor.Phone,
		Action:   actionCreateGroup,
		Response: createGroup,
	})

	var pushMessageToSpecialist = ""

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

		var status model.Status
		database.Datasource.DB().Where("name", valMessageToSpecialist).First(&status)

		notifMessageToSpecialist := util.ContentMessageToSpecialist(status.ValueNotif, specialist, doctor, url)
		userMessageToSpecialist := util.StatusMessageToSpecialist(status.ValueUser, specialist, doctor)
		pushMessageToSpecialist = util.PushMessageToSpecialist(status.ValuePush, specialist, doctor)

		log.Println(notifMessageToSpecialist)
		log.Println(userMessageToSpecialist)
		log.Println(pushMessageToSpecialist)

		if requestType == "web" {
			// NOTIF MESSAGE TO SPECIALIST
			zenzifaNotif, err := handler.ZenzivaSendSMS(specialist.Phone, notifMessageToSpecialist)
			if err != nil {
				return "", "", errors.New(err.Error())
			}
			// insert to zenziva
			database.Datasource.DB().Create(&model.Zenziva{
				Msisdn:   specialist.Phone,
				Action:   actionCreateNotifSP,
				Response: zenzifaNotif,
			})
		}

		// insert to transaction
		database.Datasource.DB().Create(
			&model.Transaction{
				UserID:       chat.UserID,
				ChatID:       chatId,
				SystemStatus: status.ValueSystem,
				NotifStatus:  notifMessageToSpecialist,
				UserStatus:   userMessageToSpecialist,
				PushStatus:   pushMessageToSpecialist,
			},
		)

	}

	return url, pushMessageToSpecialist, nil
}
