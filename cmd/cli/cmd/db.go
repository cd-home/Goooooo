package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	app  string
	mode string
)

var dbinitCmd = &cobra.Command{
	Use:   "dbinit",
	Short: "dbinitCmd",
	Long:  "dbinitCmd",
	Run: func(cmd *cobra.Command, args []string) {
		dbArgs := []string{
			"../../scripts/database.sh",
			vp.GetString("DB.HOST"),
			vp.GetString("DB.PORT"),
			vp.GetString("DB.USER"),
			vp.GetString("DB.PASSWD"),
			vp.GetString("DB.DATABASE"),
		}
		log.Println(dbArgs)
		res, err := exec.Command("sh", dbArgs...).Output()
		if err != nil {
			log.Println(err)
			log.Println(string(res))
		}
	},
}

func init() {
	rootCmd.AddCommand(dbinitCmd)
	dbinitCmd.Flags().StringVar(&app, "app", "", "chioce app")
	dbinitCmd.Flags().StringVar(&mode, "mode", "", "chioce mode")
	dbinitCmd.MarkFlagRequired("app")
	dbinitCmd.MarkFlagRequired("mode")
}
