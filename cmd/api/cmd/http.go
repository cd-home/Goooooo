package cmd

import (
	"github.com/spf13/cobra"
)

var (
	app     string
	mode    string
	configs []string
)

var httpCmd = &cobra.Command{
	Use:   "server",
	Short: "server",
	Long:  "server running on the given app and mode",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.Flags().StringVar(&app, "app", "", "chioce app")
	httpCmd.Flags().StringVar(&mode, "mode", "", "chioce mode")
	httpCmd.Flags().StringArrayVar(&configs, "config", nil, "configs")
	httpCmd.MarkFlagRequired("app")
	httpCmd.MarkFlagRequired("mode")
}
