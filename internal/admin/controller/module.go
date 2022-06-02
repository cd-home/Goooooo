package controller

import (
	v1 "github.com/GodYao1995/Goooooo/internal/admin/controller/v1"
	v2 "github.com/GodYao1995/Goooooo/internal/admin/controller/v2"
	"go.uber.org/fx"
)

var ModuleV1 = fx.Invoke(
	v1.NewUserController,
	v1.NewSysController,
	v1.NewDirectoryController,
	v1.NewJobController,
	v1.NewFileController,
)

var ModuleV2 = fx.Invoke(
	v2.NewSysController,
)
