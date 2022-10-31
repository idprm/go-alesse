package route

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/idprm/go-alesse/src/controller"
	"github.com/idprm/go-alesse/src/pkg/config"
)

func Setup(app *fiber.App) {

	app.Get("/", controller.FrontHandler)

	v1 := app.Group("v1")
	v1.Post("auth", controller.AuthHandler)
	v1.Post("verify", controller.VerifyHandler)

	healthcenters := v1.Group("healthcenters")
	healthcenters.Get("/", controller.GetAllHealthCenter)
	healthcenters.Get("/:code", controller.GetHealthCenter)

	diseases := v1.Group("diseases")
	diseases.Get("/", controller.GetAllDisease)

	doctors := v1.Group("doctors")
	doctors.Get("/", controller.GetAllDoctor)
	doctors.Get("/specialists", controller.GetAllDoctorSpecialist)
	doctors.Get("/detail/:username", controller.GetDoctor)

	medicals := v1.Group("medicals")
	medicals.Get("/", controller.GetAllMedical)
	medicals.Get("/:id", controller.GetMedical)

	chat := v1.Group("chat")
	chat.Post("/doctor", controller.ChatDoctor)
	chat.Delete("/leave", controller.ChatLeave)
	chat.Delete("/delete", controller.ChatDelete)

	/**
	 * AUTHENTICATED ROUTES
	 */
	authenticated := v1.Group("authenticated")
	authenticated.Use(jwtware.New(jwtware.Config{SigningKey: []byte(config.ViperEnv("JWT_SECRET_AUTH"))}))
	authenticated.Post("orders", controller.OrderChat)
	authenticated.Post("chat/user", controller.ChatUser)

}
