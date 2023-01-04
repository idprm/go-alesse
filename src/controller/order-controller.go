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

const (
	actionCheckUser     = "GET/CHECK USER"
	actionCreateUser    = "AUTO CREATE USER"
	actionCheckGroup    = "GET/CHECK GROUP CHANNEL"
	actionDeleteGroup   = "AUTO DELETE GROUP CHANNEL"
	actionCreateGroup   = "AUTO CREATE GROUP CHANNEL"
	actionCreateNotif   = "AUTO CREATE NOTIF TO DOCTOR"
	actionCreateMessage = "AUTO CREATE MESSAGE DOCTOR"
	actionCreateNotifSP = "AUTO CREATE NOTIF TO SPECIALIST"
	valMessageToDoctor  = "MESSAGE_TO_DOCTOR"
)

type OrderRequest struct {
	RequestType string `json:"request_type"`
	Msisdn      string `json:"msisdn"`
	DoctorID    int    `json:"doctor_id"`
}

func OrderChat(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(OrderRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var user model.User
	database.Datasource.DB().Where("msisdn", req.Msisdn).Preload("Healthcenter").First(&user)

	var doctor model.Doctor
	database.Datasource.DB().Where("id", req.DoctorID).Preload("Healthcenter").First(&doctor)

	/**
	 * Check Order
	 */

	finishUrl := config.ViperEnv("APP_HOST") + "/chat"

	/**
	 * INSERT TO ORDER
	 */
	order := model.Order{
		HealthcenterID: user.Healthcenter.ID,
		UserID:         user.ID,
		DoctorID:       doctor.ID,
		Number:         "ORD-" + util.TimeStamp(),
		Total:          0,
	}
	database.Datasource.DB().Create(&order)

	/**
	 * SENDBIRD PROCESS
	 */
	channelUrl, err := sendbirdProcess(order.ID, user.Healthcenter.ID, user.ID, doctor.ID, req.RequestType)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"status":  fiber.StatusBadGateway,
		})
	}

	var chat model.Chat
	database.Datasource.DB().Where("channel_url", channelUrl).Preload("Doctor").Preload("User").Preload("Healthcenter").First(&chat)

	var status model.Status
	database.Datasource.DB().Where("name", valMessageToDoctor).First(&status)

	if req.RequestType == "mobile" {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":        false,
			"message":      "Created Successfull",
			"status":       fiber.StatusCreated,
			"push_message": status.ValuePush,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":        false,
		"message":      "Created Successfull",
		"redirect_url": finishUrl,
		"status":       fiber.StatusCreated,
	})
}

func sendbirdProcess(orderId uint64, healthcenterId uint, userId uint64, doctorId uint, requestType string) (string, error) {

	var user model.User
	database.Datasource.DB().Where("id", userId).First(&user)

	var doctor model.Doctor
	database.Datasource.DB().Where("id", doctorId).First(&doctor)

	/**
	 * Check User Sendbird
	 */
	getUser, isUser, err := handler.SendbirdGetUser(user)
	if err != nil {
		return "", errors.New(err.Error())
	}

	/**
	 * Add User Sendbird
	 */
	database.Datasource.DB().Create(&model.Sendbird{
		Msisdn:   user.Msisdn,
		Action:   actionCheckUser,
		Response: getUser,
	})

	if isUser == true {

		// create user sendbird
		createUser, err := handler.SendbirdCreateUser(user)
		if err != nil {
			return "", errors.New(err.Error())
		}
		// insert to sendbird
		database.Datasource.DB().Create(&model.Sendbird{
			Msisdn:   user.Msisdn,
			Action:   actionCreateUser,
			Response: createUser,
		})
	}

	var chat model.Chat
	resultChat := database.Datasource.DB().Where("user_id", user.ID).First(&chat)

	getChannel, isChannel, err := handler.SendbirdGetGroupChannel(chat)
	if err != nil {
		return "", errors.New(err.Error())
	}

	// insert to sendbird
	database.Datasource.DB().Create(&model.Sendbird{
		Msisdn:   user.Msisdn,
		Action:   actionCheckGroup,
		Response: getChannel,
	})

	// check db if exist
	if resultChat.RowsAffected > 0 {
		// check channel if exist
		if isChannel == false {
			// delete channel sendbird
			deleteGroupChannel, err := handler.SendbirdDeleteGroupChannel(chat)
			if err != nil {
				return "", errors.New(err.Error())
			}
			// delete chat
			database.Datasource.DB().Where("user_id", user.ID).Delete(&chat)
			// insert to sendbirds
			database.Datasource.DB().Create(&model.Sendbird{
				Msisdn:   user.Msisdn,
				Action:   actionDeleteGroup,
				Response: deleteGroupChannel,
			})
		}
	}

	// create group
	createGroup, name, url, err := handler.SendbirdCreateGroupChannel(doctor, user)
	if err != nil {
		return "", errors.New(err.Error())
	}
	// insert to sendbird
	database.Datasource.DB().Create(&model.Sendbird{
		Msisdn:   user.Msisdn,
		Action:   actionCreateGroup,
		Response: createGroup,
	})

	if name != "" && url != "" {
		// insert to chat
		chat := model.Chat{
			HealthcenterID: healthcenterId,
			OrderID:        orderId,
			DoctorID:       doctorId,
			UserID:         userId,
			ChannelName:    name,
			ChannelUrl:     url,
		}
		database.Datasource.DB().Create(&chat)

		var status model.Status
		database.Datasource.DB().Where("name", valMessageToDoctor).First(&status)

		notifMessage := util.ContentMessageToDoctor(status.ValueNotif, user, doctor, url)
		userMessage := util.StatusMessageToDoctor(status.ValueUser, user, doctor)
		pushMessage := util.PushMessageToDoctor(status.ValuePush, user, doctor)

		log.Println(notifMessage)
		log.Println(userMessage)
		log.Println(pushMessage)

		if requestType == "web" {
			// NOTIF MESSAGE TO DOCTOR
			zenzifaNotif, err := handler.ZenzivaSendSMS(doctor.Phone, notifMessage)
			if err != nil {
				return "", errors.New(err.Error())
			}
			// insert to zenziva
			database.Datasource.DB().Create(&model.Zenziva{
				Msisdn:   doctor.Phone,
				Action:   actionCreateNotif,
				Response: zenzifaNotif,
			})
		}

		// auto message to user
		autoMessageDoctor, err := handler.SendbirdAutoMessageDoctor(url, doctor, user)
		if err != nil {
			return "", errors.New(err.Error())
		}

		// insert to sendbird
		database.Datasource.DB().Create(&model.Sendbird{
			Msisdn:   user.Msisdn,
			Action:   actionCreateMessage,
			Response: autoMessageDoctor,
		})

		// insert to transaction
		database.Datasource.DB().Create(
			&model.Transaction{
				UserID:       chat.UserID,
				ChatID:       chat.ID,
				SystemStatus: status.ValueSystem,
				NotifStatus:  notifMessage,
				UserStatus:   userMessage,
				PushStatus:   pushMessage,
			},
		)

	}

	return url, nil
}
