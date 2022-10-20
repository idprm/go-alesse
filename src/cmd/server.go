package cmd

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/idprm/go-alesse/src/pkg/route"
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

		secret, err := localconfig.LoadSecret("src/server/secret.yaml")
		if err != nil {
			panic(err)
		}

		app := fiber.New()

		app.Use(cors.New())

		route.Setup(secret, app)

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
