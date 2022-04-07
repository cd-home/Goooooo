package controller

import (
	v1 "github.com/GodYao1995/Goooooo/internal/api/controller/v1"
	"go.uber.org/fx"
)

var Module = fx.Invoke(
	v1.NewOrderController,
)
