package logic

import (
	v1 "github.com/GodYao1995/Goooooo/internal/admin/logic/v1"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	v1.NewUserLogic,
	v1.NewDirectoryrLogic,
	v1.NewFileLogic,
)
