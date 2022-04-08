package config

import (
	"flag"
	"fmt"
	"strings"
)

func command() (string, string, ConfigPaths) {
	var app string
	var mode string
	var configPaths ConfigPaths
	flag.StringVar(&app, "app", "", "应用")
	flag.StringVar(&mode, "mode", "", "运行环境")
	flag.Var(&configPaths, "config", "配置文件目录")
	flag.Parse()
	return app, mode, configPaths
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
