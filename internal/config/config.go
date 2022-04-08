package config

import (
	"log"
	"os"

	"github.com/GodYao1995/Goooooo/pkg/config"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewConfig)

func _env() (string, string) {
	return os.Getenv("APP_NAME"), os.Getenv("APP_MODE")
}

func NewConfig() *viper.Viper {
	// Development or Test environments
	app, mode, configPaths := command()

	// Production Mode Use the environment variables
	if _app, _mode := _env(); _app != "" && _mode != "" {
		app = _app
		mode = _mode
	}

	if app == "" && mode == "" {
		log.Println("未找到应用与环境信息")
		os.Exit(1)
	}
	return config.NewViper(app, mode, configPaths.GetNames()...)
}
