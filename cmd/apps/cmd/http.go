package cmd

import (
	"github.com/spf13/cobra"
)

var (
	app  string
	mode string
)

var apiHttpCmd = &cobra.Command{
	Use:   "api",
	Short: "api",
	Long:  "api server running on the given app and mode",
	Run: func(cmd *cobra.Command, args []string) {
		RunApi()
	},
}

var adminHttpCmd = &cobra.Command{
	Use:   "admin",
	Short: "admin",
	Long:  "admin server running on the given app and mode",
	Run: func(cmd *cobra.Command, args []string) {
		RunAdmin()
	},
}

func init() {
	rootCmd.AddCommand(apiHttpCmd, adminHttpCmd)
	apiHttpCmd.Flags().StringVar(&app, "app", "", "chioce app")
	apiHttpCmd.Flags().StringVar(&mode, "mode", "", "chioce mode")
	apiHttpCmd.MarkFlagRequired("app")
	apiHttpCmd.MarkFlagRequired("mode")

	adminHttpCmd.Flags().StringVar(&app, "app", "", "chioce app")
	adminHttpCmd.Flags().StringVar(&mode, "mode", "", "chioce mode")
	adminHttpCmd.MarkFlagRequired("app")
	adminHttpCmd.MarkFlagRequired("mode")
}
