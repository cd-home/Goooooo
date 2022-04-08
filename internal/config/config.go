package config

import (
	"log"
	"os"
	"strings"

	"github.com/GodYao1995/Goooooo/pkg/config"
	"github.com/GodYao1995/Goooooo/pkg/errno"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewConfig)

func NewConfig() *viper.Viper {
	// Development or Test environments
	app, mode, configPaths := command()

	// Production Mode Use the environment variables
	_differentAppAndEnvironment(&app, &mode)

	return config.NewViper(app, mode, configPaths.GetNames()...)
}

func _differentAppAndEnvironment(app, mode *string) {
	_app, _mode := os.Getenv("APP_NAME"), os.Getenv("APP_MODE")
	if len(strings.TrimSpace(_app)) != 0 && len(strings.TrimSpace(_mode)) != 0 {
		*app = _app
		*mode = _mode
	}

	if len(strings.TrimSpace(*app)) == 0 || len(strings.TrimSpace(*mode)) == 0 {
		log.Fatalln(errno.NotFoundService)
	}
}
