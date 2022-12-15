package cmd

import (
	"log"
	"time"

	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
	"github.com/spf13/cobra"
)

var callbackCmd = &cobra.Command{
	Use:   "callback",
	Short: "Callback CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		rows, err := database.Datasource.DB().Model(&model.Chat{}).Where("is_leave", false).Where("created_at < NOW() - INTERVAL 6 HOUR").Rows()
		if err != nil {
			log.Println(err)
		}
		defer rows.Close()

		for rows.Next() {
			var ch model.Chat
			database.Datasource.DB().ScanRows(rows, &ch)

			var chat model.Chat
			database.Datasource.DB().Where("id", ch.ID).Preload("Order").Preload("User").Preload("Doctor").First(&chat)

			// update chat is leave = true
			database.Datasource.DB().Model(&model.Chat{}).Where("id", ch.ID).Updates(&model.Chat{IsLeave: true, LeaveAt: time.Now()})

		}
	},
}

func init() {
	rootCmd.AddCommand(callbackCmd)
}
