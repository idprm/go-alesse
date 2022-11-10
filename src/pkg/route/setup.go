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
	auth := v1.Group("auth")

	auth.Post("login", controller.LoginHandler)
	auth.Post("verify", controller.VerifyHandler)
	auth.Post("register", controller.RegisterHandler)

	healthcenters := v1.Group("healthcenters")
	healthcenters.Get("/", controller.GetAllHealthCenter)
	healthcenters.Get("/:code", controller.GetHealthCenter)

	diseases := v1.Group("diseases")
	diseases.Get("/", controller.GetAllDisease)

	medicines := v1.Group("medicines")
	medicines.Get("/", controller.GetAllMedicine)
	medicines.Get("/:id", controller.GetMedicine)

	doctors := v1.Group("doctors")
	doctors.Get("/", controller.GetAllDoctor)
	doctors.Get("/detail/:username", controller.GetDoctor)

	specialists := v1.Group("specialists")
	specialists.Get("/", controller.GetAllSpecialist)
	specialists.Get("/detail/:username", controller.GetSpecialist)

	officers := v1.Group("officers")
	officers.Get("/type/:type", controller.GetAllOfficer)
	officers.Get("/phone/:phone", controller.GetOfficer)

	chat := v1.Group("chat")
	chat.Post("/doctor", controller.ChatDoctor)
	chat.Delete("/leave", controller.ChatLeave)
	chat.Delete("/delete", controller.ChatDelete)

	medicalresumes := v1.Group("medicalresumes")
	medicalresumes.Get("/", controller.GetAllMedicalResume)
	medicalresumes.Get("/detail/:slug", controller.GetMedicalResume)
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
	homecares.Get("/doctor/:slug", controller.GetHomecareByDoctor)
	homecares.Post("/", controller.SaveHomecare)
	homecares.Get("/photos/:id", controller.GetHomecareAllPhoto)
	homecares.Post("/photos", controller.UploadPhoto)
	homecares.Post("/officer", controller.SaveHomecareOfficer)

	/**
	 * AUTHENTICATED ROUTES
	 */
	authenticated := v1.Group("authenticated")
	authenticated.Use(jwtware.New(jwtware.Config{SigningKey: []byte(config.ViperEnv("JWT_SECRET_AUTH"))}))
	authenticated.Post("orders", controller.OrderChat)
	authenticated.Post("chat/user", controller.ChatUser)

}
