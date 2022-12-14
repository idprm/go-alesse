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
	medicines.Get("/healthcenter/:healthcenter", controller.GetMedicineByHealthCenter)
	medicines.Get("/:id", controller.GetMedicine)

	doctors := v1.Group("doctors")
	doctors.Get("/", controller.GetAllDoctor)
	doctors.Get("/healthcenter/:healthcenter", controller.GetDoctorByHealthCenter)
	doctors.Get("/detail/:username", controller.GetDoctor)

	specialists := v1.Group("specialists")
	specialists.Get("/", controller.GetAllSpecialist)
	specialists.Get("/healthcenter/:healthcenter", controller.GetSpecialistByHealthCenter)
	specialists.Get("/detail/:username", controller.GetSpecialist)

	officers := v1.Group("officers")
	officers.Get("/", controller.GetAllOfficer)
	officers.Get("/healthcenter/:healthcenter", controller.GetOfficerByHealthCenter)
	officers.Get("/:phone", controller.GetOfficer)

	apothecaries := v1.Group("apothecaries")
	apothecaries.Get("/", controller.GetAllApothecary)
	apothecaries.Get("/healthcenter/:healthcenter", controller.GetApothecaryByHealthCenter)
	apothecaries.Get("/:phone", controller.GetApothecary)

	couriers := v1.Group("couriers")
	couriers.Get("/", controller.GetAllCourier)
	couriers.Get("/healthcenter/:healthcenter", controller.GetCourierByHealthCenter)
	couriers.Get("/:phone", controller.GetCourier)

	drivers := v1.Group("drivers")
	drivers.Get("/", controller.GetAllDriver)
	drivers.Get("/healthcenter/:healthcenter", controller.GetDriverByHealthCenter)
	drivers.Get("/:phone", controller.GetDriver)

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
	prescriptions.Get("/medicines/:slug", controller.GetAllPrescriptionMedicine)
	prescriptions.Post("/medicines", controller.SavePrescription)

	homecares := v1.Group("homecares")
	homecares.Get("/", controller.GetAllHomecare)
	homecares.Get("/:slug", controller.GetHomecare)
	homecares.Get("/doctor/:slug", controller.GetHomecareByDoctor)
	homecares.Get("/medicines/:slug", controller.GetAllHomecareByMedicines)
	homecares.Post("/", controller.SaveHomecare)
	homecares.Get("/photos/:id", controller.GetHomecareAllPhoto)
	homecares.Post("/photos", controller.HomecarePhoto)
	homecares.Get("/officer/:slug", controller.GetHomecareByOfficer)
	homecares.Post("/officer", controller.SaveHomecareOfficer)
	homecares.Post("/resume", controller.SaveHomecareResume)

	pharmacies := v1.Group("pharmacies")
	pharmacies.Get("/", controller.GetAllPharmacy)
	pharmacies.Get("/:slug", controller.GetPharmacy)
	pharmacies.Get("/doctor/:slug", controller.GetPharmacyByDoctor)
	pharmacies.Get("/medicines/:slug", controller.GetAllPharmacyByMedicines)
	pharmacies.Get("/apothecary/:slug", controller.GetPharmacyByApothecary)
	pharmacies.Get("/courier/:slug", controller.GetPharmacyByCourier)
	pharmacies.Post("/", controller.SavePharmacy)
	pharmacies.Post("/process", controller.SaveProcessPharmacy)
	pharmacies.Post("/sent", controller.SaveSentPharmacy)
	pharmacies.Post("/take", controller.SaveTakePharmacy)
	pharmacies.Post("/finish", controller.SaveFinishPharmacy)
	pharmacies.Get("/photos/:id", controller.GetPharmacyAllPhoto)
	pharmacies.Post("/photos", controller.PharmacyPhoto)

	referrals := v1.Group("referrals")
	referrals.Post("/", controller.Referral)
	referrals.Get("/chat", controller.ReferralChat)

	feedbacks := v1.Group("feedbacks")
	feedbacks.Get("/:slug", controller.GetFeedback)
	feedbacks.Post("/", controller.SaveFeedback)

	users := v1.Group("users")
	users.Get("/:msisdn", controller.GetUserByMsisdn)
	users.Get("/transactions/:msisdn", controller.GetTransactionByUser)

	histories := v1.Group("histories")
	histories.Get("/:msisdn", controller.GetMedicalHistory)

	channels := v1.Group("channels")
	channels.Get("/doctors/:channel", controller.GetDoctorByChannel)
	channels.Get("/apothecaries/:channel", controller.GetApothecaryByChannel)
	channels.Get("/officers/:channel", controller.GetOfficerByChannel)
	channels.Get("/drivers/:channel", controller.GetDriverByChannel)

	/**
	 * AUTHENTICATED ROUTES
	 */
	authenticated := v1.Group("authenticated")
	authenticated.Use(jwtware.New(jwtware.Config{SigningKey: []byte(config.ViperEnv("JWT_SECRET_AUTH"))}))
	authenticated.Post("orders", controller.OrderChat)
	authenticated.Post("chat/user", controller.ChatUser)

	/**
	 * version 2 for mobile
	 */
	v2 := app.Group("v2")
	auth2 := v2.Group("auth")

	auth2.Post("login", controller.MLoginHandler)
	auth2.Post("register", controller.MRegisterHandler)
	auth2.Post("verify", controller.MVerifyHandler)

	auth2.Post("doctor", controller.MAuthDoctorHandler)
	auth2.Post("officer", controller.MAuthOfficerHandler)
	auth2.Post("apothecary", controller.MAuthApothecaryHandler)
	auth2.Post("courier", controller.MAuthCourierHandler)
	auth2.Post("specialist", controller.MAuthSpecialistHandler)

}
