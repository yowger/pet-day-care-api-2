package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yowger/pet-day-care-api-2/bootstrap"
)

func init() {
	rootCmd.AddCommand(appCmd)
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Starts the Pet Day Care API application",
	Long:  "Database seeder",
	Run: func(cmd *cobra.Command, args []string) {
		startApp()
	},
}

func startApp() {
	bootstrap.App()
}
