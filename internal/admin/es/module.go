package xes

import (
	v1 "github.com/cd-home/Goooooo/internal/admin/es/v1"
	"go.uber.org/fx"
)

var Module = fx.Provide(v1.NewUserEs)
