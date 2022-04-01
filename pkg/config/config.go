package config

import (
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
	_DefaultConfigType = "toml"
)

func defaultConfigPath(mode string, configPaths ...string) []string {
	if len(configPaths) == 0 {
		return []string{_DefaultConfigDot, _DefaultConfigCur + mode, _DefaultConfigPar + mode}
	}
	var paths []string
	for _, path := range configPaths {
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
		paths = append(paths, path+mode)
	}
	return paths
}

func NewViper(app string, mode string, configPaths ...string) *viper.Viper {
	vp := viper.New()
	// Development, Testing mode is always Read FileConfig
	if mode == _Development || mode == _ShortDevelopment || mode == _Testing || mode == _ShortTesting {
		vp.SetConfigName(app)
		vp.SetConfigType(_DefaultConfigType)
		_configPaths := defaultConfigPath(mode, configPaths...)
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
