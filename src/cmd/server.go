package cmd

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/handler"
	"github.com/idprm/go-alesse/src/pkg/util/localconfig"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		config, err := localconfig.LoadConfig("src/server/config.yaml")
		if err != nil {
			panic(err)
		}

		app := fiber.New()

		// version 1
		v1 := app.Group("v1")

		/**
		 * FRONTEND ROUTES
		 */
		v1.Post("register", handler.Register)
		v1.Post("login", handler.Login)
		v1.Post("verify", handler.Verify)

		auth := v1.Group("auth")
		auth.Get("chat", handler.GetChat)
		auth.Get("medical", handler.GetMedical)

		path, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		app.Static("/", path+"/public")

		log.Fatal(app.Listen(":" + config.APP.Port))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
