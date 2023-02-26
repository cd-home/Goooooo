package repository

import (
	v1 "github.com/cd-home/Goooooo/internal/admin/repository/v1"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	v1.NewUserRepository,
	v1.NewDirectoryRepository,
	v1.NewFileRepository,
	v1.NewRoleRepository,
)
