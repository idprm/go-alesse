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
)

type OrderRequest struct {
	Msisdn   string `json:"msisdn"`
	DoctorID int    `json:"doctor_id"`
}

type ReferralRequest struct {
	DoctorID     int `json:"doctor_id"`
	SpecialistID int `json:"specialist_id"`
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
	database.Datasource.DB().Where("msisdn", req.Msisdn).First(&user)

	var doctor model.Doctor
	database.Datasource.DB().Where("id", req.DoctorID).First(&doctor)

	/**
	 * Check Order
	 */
	var order model.Order
	resultOrder := database.Datasource.DB().
		Where("user_id", user.ID).
		Where("doctor_id", doctor.ID).
		Where("DATE(created_at) = DATE(?)", time.Now()).
		First(&order)

	finishUrl := config.ViperEnv("APP_HOST") + "/chat"

	if resultOrder.RowsAffected == 0 {
		/**
		 * INSERT TO ORDER
		 */
		database.Datasource.DB().Create(&model.Order{
			UserID:   user.ID,
			DoctorID: doctor.ID,
			Number:   "ORD-" + util.TimeStamp(),
			Total:    0,
		})

		/**
		 * SENDBIRD PROCESS
		 */
		err := sendbirdProcess(user.ID, doctor.ID)
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

	req := new(ReferralRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var doctor model.Doctor
	database.Datasource.DB().Where("id", req.DoctorID).Where("is_specialist", false).First(&doctor)

	var specialist model.Doctor
	database.Datasource.DB().Where("id", req.SpecialistID).Where("is_specialist", true).First(&specialist)

	/**
	 * Check Order
	 */
	var referral model.Referral
	resultReferral := database.Datasource.DB().
		Where("doctor_id", doctor.ID).
		Where("doctor_specialist_id", specialist.ID).
		Where("DATE(created_at) = DATE(?)", time.Now()).
		First(&referral)

	finishUrl := config.ViperEnv("APP_HOST") + "/specialist/chat"

	if resultReferral.RowsAffected == 0 {
		/**
		 * INSERT TO Referral
		 */
		database.Datasource.DB().Create(&model.Referral{
			SpecialistID: specialist.ID,
			DoctorID:     doctor.ID,
		})

		/**
		 * SENDBIRD PROCESS
		 */
		// err := sendbirdProcess(user.ID, doctor.ID)
		// if err != nil {
		// 	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
		// 		"error":   true,
		// 		"message": err.Error(),
		// 		"status":  fiber.StatusBadGateway,
		// 	})
		// }

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

func sendbirdProcess(userId uint64, doctorId uint) error {

	var order model.Order
	database.Datasource.DB().
		Where("user_id", userId).
		Where("doctor_id", doctorId).
		Where("DATE(created_at) = DATE(?)", time.Now()).
		Preload("User").Preload("Doctor").
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
		database.Datasource.DB().Create(&model.Chat{
			OrderID:     order.ID,
			DoctorID:    order.Doctor.ID,
			UserID:      order.User.ID,
			ChannelName: name,
			ChannelUrl:  url,
		})

		var conf model.Config
		database.Datasource.DB().Where("name", "NOTIF_MESSAGE_DOCTOR").First(&conf)

		urlWeb := config.ViperEnv("APP_HOST") + "/chat/" + url
		replaceMessage := strings.NewReplacer("@v1", order.Doctor.Name, "@v2", order.User.Name, "@v3", urlWeb)
		message := replaceMessage.Replace(conf.Value)

		// NOTIF MESSAGE TO DOCTOR
		zenzifaNotif, err := handler.ZenzivaSendSMS(order.Doctor.Phone, message)
		if err != nil {
			return errors.New(err.Error())
		}
		// insert to zenziva
		database.Datasource.DB().Create(&model.Zenziva{
			Msisdn:   order.User.Msisdn,
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
	}

	return nil
}

func sendbirdProcessReferral(specialistId uint, doctorId uint) error {

	var referral model.Referral
	database.Datasource.DB().
		Where("doctor_specialist_id", specialistId).
		Where("doctor_id", doctorId).
		Where("DATE(created_at) = DATE(?)", time.Now()).
		Preload("Doctor").
		Preload("Specialist").
		First(&referral)

		/**
		 * Check User Sendbird
		 */
	getSpecialist, isSpecialist, err := handler.SendbirdGetSpecialist(referral.Specialist)
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

	if isSpecialist == true {

		// create user sendbird
		createUser, err := handler.SendbirdCreateSpecialist(referral.Specialist)
		if err != nil {
			return errors.New(err.Error())
		}
		// insert to sendbird
		database.Datasource.DB().Create(&model.Sendbird{
			Msisdn:   referral.Specialist.Phone,
			Action:   actionCreateUser,
			Response: createUser,
		})
	}

	return nil
}
