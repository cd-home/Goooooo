package cmd

import (
	"github.com/spf13/cobra"
)

var (
	app     string
	mode    string
	configs []string
)

var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "job",
	Long:  "job running on the given app and mode",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func init() {
	rootCmd.AddCommand(jobCmd)
	jobCmd.Flags().StringVar(&app, "app", "", "chioce app")
	jobCmd.Flags().StringVar(&mode, "mode", "", "chioce mode")
	jobCmd.Flags().StringArrayVar(&configs, "config", nil, "configs")
	jobCmd.MarkFlagRequired("app")
	jobCmd.MarkFlagRequired("mode")
}
