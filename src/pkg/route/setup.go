package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/controller"
	"github.com/idprm/go-alesse/src/handler"
)

func Setup(app *fiber.App) {
	app.Get("/", controller.FrontHandler)

	v1 := app.Group("v1")
	v1.Post("auth", controller.AuthHandler)

	v1.Post("login", handler.Login)
	v1.Post("register", handler.Register)

}
