package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

const (
	_Development      = "development"
	_Testing          = "testing"
	_Production       = "production"
	_ShortDevelopment = "dev"
	_ShortTesting     = "test"
	_ShortProduction  = "prod"
)

const (
	_DefaultConfigDot  = "."
	_DefaultConfigCur  = "./configs/"
	_DefaultConfigPar  = "../configs/"
	_DefaultConfigPPar = "../../configs/"
	_DefaultConfigType = "toml"
)

func defaultConfigPath(app string, configPaths ...string) []string {
	if len(configPaths) == 0 {
		return []string{
			_DefaultConfigDot,
			_DefaultConfigCur + app,
			_DefaultConfigPar + app,
			_DefaultConfigPPar + app,
		}
	}
	var paths []string
	for _, path := range configPaths {
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
		paths = append(paths, path+app)
	}
	return paths
}

func NewViper(app string, mode string, configPaths ...string) *viper.Viper {
	vp := viper.New()
	// Development, Testing mode is always Read FileConfig
	if mode == _Development || mode == _ShortDevelopment || mode == _Testing || mode == _ShortTesting {
		vp.SetConfigName(mode)
		vp.SetConfigType(_DefaultConfigType)
		_configPaths := defaultConfigPath(app, configPaths...)
		log.Println(_configPaths)
		for _, path := range _configPaths {
			vp.AddConfigPath(path)
		}
		if err := vp.ReadInConfig(); err != nil {
			panic(err)
		}
		// Production mode is always Read Env
	} else {
		vp.SetEnvPrefix(strings.ToUpper(app))
		vp.AutomaticEnv()
	}
	return vp
}
