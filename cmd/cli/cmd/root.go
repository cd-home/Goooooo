package cmd

import (
	"fmt"

	"github.com/GodYao1995/Goooooo/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var vp *viper.Viper

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "init scripts",
	Long:  "init db scripts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rootCmd")
	},
}

func Exeute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	vp = config.NewViper(app, mode)
}
