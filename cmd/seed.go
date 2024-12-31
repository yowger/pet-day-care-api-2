package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Database seeder",
	Run: func(cmd *cobra.Command, args []string) {
		seedDB()
	},
}

func seedDB() {
	fmt.Println("seeding...")
}
