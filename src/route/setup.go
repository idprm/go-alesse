package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/controller"
)

func Setup(app *fiber.App) {

	// version 1
	v1 := app.Group("v1")

	/**
	 * FRONTEND ROUTES
	 */
	v1.Post("register", controller.Register)
	v1.Post("login", controller.Login)
	v1.Post("verify", controller.Verify)

	auth := v1.Group("auth")
	auth.Get("chat", controller.GetChat)
	auth.Get("medical", controller.GetMedical)

}
