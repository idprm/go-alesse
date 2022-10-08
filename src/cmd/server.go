package cmd

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/config"
	"github.com/idprm/go-alesse/src/route"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		app := fiber.New()

		route.Setup(app)

		path, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		app.Static("/", path+"/public")

		var conf config.APPConfig
		log.Fatal(app.Listen(":" + conf.GetAPPPort()))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
