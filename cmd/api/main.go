package main

import (
	_ "github.com/cd-home/Goooooo/api/api"
	"github.com/cd-home/Goooooo/cmd/api/cmd"
)

// InitRouter @title Goooooo-Api
// @contact.name God Yao
// @contact.email liyaoo1995@163.com
// @version 1.0.0
// @description this is Goooooo-Api Sys.
// @host 127.0.0.1:8081
// @BasePath /api/v1
func main() {
	cmd.Exeute()
}
