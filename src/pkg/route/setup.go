package route

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/idprm/go-alesse/src/controller"
	"github.com/idprm/go-alesse/src/pkg/util/localconfig"
)

func Setup(cfg *localconfig.Secret, app *fiber.App) {

	app.Get("/", controller.FrontHandler)

	v1 := app.Group("v1")
	v1.Post("auth", controller.AuthHandler)

	doctor := v1.Group("doctors")
	doctor.Get("/", controller.GetAllDoctor)
	doctor.Get("/:username", controller.GetDoctor)

	medicals := v1.Group("medicals")
	medicals.Get("/", controller.GetAllMedical)
	medicals.Get("/:id", controller.GetMedical)

	chat := v1.Group("chat")
	chat.Post("/doctor", controller.ChatDoctor)
	// chat.Delete("/leave", controller.ChatLeave)
	// chat.Delete("/delete", controller.ChatDelete)

	/**
	 * AUTHENTICATED ROUTES
	 */
	authenticated := v1.Group("authenticated")
	authenticated.Use(jwtware.New(jwtware.Config{SigningKey: []byte(cfg.JWT.Secret)}))
	// authenticated.Post("orders", controller.OrderChat)
	// authenticated.Post("chat/user", controller.ChatUser)

}
