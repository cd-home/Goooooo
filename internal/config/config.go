package config

import (
	"log"
	"os"

	"github.com/GodYao1995/Goooooo/pkg/config"
	"github.com/GodYao1995/Goooooo/pkg/errno"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewConfig)

func NewConfig() *viper.Viper {
	// Development or Test environments
	app, mode, configPaths := command()

	// Production Mode Use the environment variables
	if _app, _mode := _differentAppAndEnvironment(); _app != "" && _mode != "" {
		app = _app
		mode = _mode
	}

	if app == "" && mode == "" {
		log.Fatalln(errno.NotFoundService)
	}
	return config.NewViper(app, mode, configPaths.GetNames()...)
}

func _differentAppAndEnvironment() (string, string) {
	return os.Getenv("APP_NAME"), os.Getenv("APP_MODE")
}
