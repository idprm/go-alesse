package cmd

import (
	"context"

	"github.com/idprm/go-alesse/src/database"
	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Worker CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		for {
			processorZenziva("xxx")
		}
	},
}

func processorZenziva(key string) {

	database.NewRedisClient().Get(context.Background(), key)

}

func init() {
	rootCmd.AddCommand(workerCmd)
}
