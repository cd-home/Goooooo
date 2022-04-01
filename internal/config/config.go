package config

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/GodYao1995/Goooooo/pkg/config"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Invoke(NewConfig)

const (
	Version = "1.0.0"
	AppName = "admin"
	Mode    = "dev"
)

func NewConfig() *viper.Viper {
	var app string
	var mode string
	var configPaths ConfigPaths

	flag.StringVar(&app, "app", AppName, "应用")
	flag.StringVar(&mode, "mode", Mode, "运行环境")
	flag.Var(&configPaths, "config", "配置文件目录")

	flag.Parse()

	log.Println("app: ", app)
	log.Println("mode: ", mode)
	log.Println("config: ", configPaths.GetNames())

	return config.NewViper(app, mode, configPaths.GetNames()...)
}

type ConfigPaths struct {
	Paths []string
}

func (c *ConfigPaths) GetNames() []string {
	return c.Paths
}

func (c *ConfigPaths) String() string {
	return fmt.Sprint(c.Paths)
}

func (c *ConfigPaths) Set(v string) error {
	if len(c.Paths) > 0 {
		return fmt.Errorf("no")
	}
	paths := strings.Split(v, ",")
	c.Paths = append(c.Paths, paths...)
	return nil
}
