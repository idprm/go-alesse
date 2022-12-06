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
)

type OrderRequest struct {
	Msisdn   string `json:"msisdn"`
	DoctorID int    `json:"doctor_id"`
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
	var order model.Order
	resultOrder := database.Datasource.DB().
		Where("healthcenter_id", user.HealthcenterID).
		Where("user_id", user.ID).
		Where("doctor_id", doctor.ID).
		First(&order)

	finishUrl := config.ViperEnv("APP_HOST") + "/chat"

	if resultOrder.RowsAffected == 0 {
		/**
		 * INSERT TO ORDER
		 */
		database.Datasource.DB().Create(&model.Order{
			HealthcenterID: user.Healthcenter.ID,
			UserID:         user.ID,
			DoctorID:       doctor.ID,
			Number:         "ORD-" + util.TimeStamp(),
			Total:          0,
		})

		/**
		 * SENDBIRD PROCESS
		 */
		err := sendbirdProcess(user.Healthcenter.ID, user.ID, doctor.ID)
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

func sendbirdProcess(healthcenterId uint, userId uint64, doctorId uint) error {

	var order model.Order
	database.Datasource.DB().
		Where("healthcenter_id", healthcenterId).
		Where("user_id", userId).
		Where("doctor_id", doctorId).
		Preload("Healthcenter").Preload("User").Preload("Doctor").
		First(&order)
	/**
	 * Check User Sendbird
	 */
	getUser, isUser, err := handler.SendbirdGetUser(order.User)
	if err != nil {
		return errors.New(err.Error())
	}

	/**
	 * Add User Sendbird
	 */
	database.Datasource.DB().Create(&model.Sendbird{
		Msisdn:   order.User.Msisdn,
		Action:   actionCheckUser,
		Response: getUser,
	})

	if isUser == true {

		// create user sendbird
		createUser, err := handler.SendbirdCreateUser(order.User)
		if err != nil {
			return errors.New(err.Error())
		}
		// insert to sendbird
		database.Datasource.DB().Create(&model.Sendbird{
			Msisdn:   order.User.Msisdn,
			Action:   actionCreateUser,
			Response: createUser,
		})
	}

	var chat model.Chat
	resultChat := database.Datasource.DB().Where("user_id", order.User.ID).First(&chat)

	getChannel, isChannel, err := handler.SendbirdGetGroupChannel(chat)
	if err != nil {
		return errors.New(err.Error())
	}

	// insert to sendbird
	database.Datasource.DB().Create(&model.Sendbird{
		Msisdn:   order.User.Msisdn,
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
				return errors.New(err.Error())
			}
			// delete chat
			database.Datasource.DB().Where("user_id", order.User.ID).Delete(&chat)
			// insert to sendbirds
			database.Datasource.DB().Create(&model.Sendbird{
				Msisdn:   order.User.Msisdn,
				Action:   actionDeleteGroup,
				Response: deleteGroupChannel,
			})
		}
	}

	// create group
	createGroup, name, url, err := handler.SendbirdCreateGroupChannel(order.Doctor, order.User)
	if err != nil {
		return errors.New(err.Error())
	}
	// insert to sendbird
	database.Datasource.DB().Create(&model.Sendbird{
		Msisdn:   order.User.Msisdn,
		Action:   actionCreateGroup,
		Response: createGroup,
	})

	if name != "" && url != "" {
		// insert to chat
		chat := model.Chat{
			HealthcenterID: order.HealthcenterID,
			OrderID:        order.ID,
			DoctorID:       order.Doctor.ID,
			UserID:         order.User.ID,
			ChannelName:    name,
			ChannelUrl:     url,
		}
		database.Datasource.DB().Create(&chat)

		const (
			valMessageToDoctor = "MESSAGE_TO_DOCTOR"
		)

		var status model.Status
		database.Datasource.DB().Where("name", valMessageToDoctor).First(&status)

		notifMessage := util.ContentMessageToDoctor(status.ValueNotif, order, url)
		userMessage := util.StatusMessageToDoctor(status.ValueUser, order)

		log.Println(notifMessage)
		log.Println(userMessage)

		// NOTIF MESSAGE TO DOCTOR
		zenzifaNotif, err := handler.ZenzivaSendSMS(order.Doctor.Phone, notifMessage)
		if err != nil {
			return errors.New(err.Error())
		}
		// insert to zenziva
		database.Datasource.DB().Create(&model.Zenziva{
			Msisdn:   order.Doctor.Phone,
			Action:   actionCreateNotif,
			Response: zenzifaNotif,
		})

		// auto message to user
		autoMessageDoctor, err := handler.SendbirdAutoMessageDoctor(url, order.Doctor, order.User)
		if err != nil {
			return errors.New(err.Error())
		}

		// insert to sendbird
		database.Datasource.DB().Create(&model.Sendbird{
			Msisdn:   order.User.Msisdn,
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
			},
		)

	}

	return nil
}
