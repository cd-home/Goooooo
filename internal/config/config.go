package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/GodYao1995/Goooooo/pkg/config"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewConfig)

const (
	MockProd = false
)

func _env() (string, string) {
	return os.Getenv("APP"), os.Getenv("MODE")
}

func NewConfig() *viper.Viper {
	var app string
	var mode string
	var configPaths ConfigPaths

	flag.StringVar(&app, "app", "", "应用")
	flag.StringVar(&mode, "mode", "", "运行环境")
	flag.Var(&configPaths, "config", "配置文件目录")

	flag.Parse()

	// Production Mode
	if _app, _mode := _env(); _app != "" && _mode != "" {
		app = _app
		mode = _mode
	}

	if app == "" && mode == "" {
		log.Println("未找到应用与环境信息")
		os.Exit(1)
	}

	log.Println("app: ", app)
	log.Println("mode: ", mode)
	log.Println("config: ", configPaths.GetNames())

	vp := config.NewViper(app, mode, configPaths.GetNames()...)

	// common Get mode
	log.Println(vp.GetString(app + ".DB_URL"))
	log.Println(vp.GetString(app + ".SECRET"))

	vp.SetDefault("APP", app)
	vp.SetDefault("MODE", mode)

	log.Println(vp.GetString(app + ".SERVER_HOST"))

	return vp
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
