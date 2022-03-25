package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const (
	_Development      = "development"
	_Testing          = "testing"
	_PRODUCTION       = "production"
	_ShortDevelopment = "dev"
	_ShortTesting     = "test"
	_ShortProduction  = "prod"
)

const (
	_DefaultConfigType = "toml"
)

func defaultConfigPath(configPaths ...string) []string {
	if len(configPaths) == 0 {
		return []string{".", "./configs", "../configs"}
	}
	return configPaths
}

func NewViper(mode string, configPaths ...string) *viper.Viper {
	vp := viper.New()
	// Development mode is always Read FileConfig
	if mode == _Development || mode == _ShortDevelopment {
		vp.SetConfigName(mode)
		vp.SetConfigType(_DefaultConfigType)
		_configPaths := defaultConfigPath(configPaths...)
		for _, path := range _configPaths {
			vp.AddConfigPath(path)
		}
		if err := vp.ReadInConfig(); err != nil {
			log.Println(err.Error())
			panic(err)
		}
		// Production mode is always Read Env
	} else {
		vp.AutomaticEnv()
	}
	return vp
}
