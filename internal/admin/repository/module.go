package repository

import (
	v1 "github.com/GodYao1995/Goooooo/internal/admin/repository/v1"
	"go.uber.org/fx"
)

var Module = fx.Provide(v1.NewUserRepository, v1.NewDirectoryRepository)
