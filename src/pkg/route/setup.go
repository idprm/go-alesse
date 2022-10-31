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
	diseases.Get("/:id")

	doctors := v1.Group("doctors")
	doctors.Get("/", controller.GetAllDoctor)
	doctors.Get("/specialists", controller.GetAllDoctorSpecialist)
	doctors.Get("/detail/:username", controller.GetDoctor)

	chat := v1.Group("chat")
	chat.Post("/doctor", controller.ChatDoctor)
	chat.Delete("/leave", controller.ChatLeave)
	chat.Delete("/delete", controller.ChatDelete)

	medicalresumes := v1.Group("medicalresumes")
	medicalresumes.Get("/", controller.GetAllMedicalResume)
	medicalresumes.Get("/detail/:number", controller.GetMedicalResume)
	medicalresumes.Post("/", controller.SaveMedicalResume)

	prescriptions := v1.Group("prescriptions")
	prescriptions.Get("/", controller.GetAllPrescription)
	prescriptions.Get("/detail/:number", controller.GetPrescription)
	prescriptions.Post("/", controller.SavePrescription)
	prescriptions.Get("/medicines", controller.GetAllPrescriptionMedicine)
	prescriptions.Post("/medicines", controller.SavePrescription)

	homecares := v1.Group("homecares")
	homecares.Get("/", controller.GetAllHomecare)
	homecares.Get("/detail/:number", controller.GetHomecare)
	homecares.Post("/", controller.SaveHomecare)

	/**
	 * AUTHENTICATED ROUTES
	 */
	authenticated := v1.Group("authenticated")
	authenticated.Use(jwtware.New(jwtware.Config{SigningKey: []byte(config.ViperEnv("JWT_SECRET_AUTH"))}))
	authenticated.Post("orders", controller.OrderChat)
	authenticated.Post("chat/user", controller.ChatUser)

}
