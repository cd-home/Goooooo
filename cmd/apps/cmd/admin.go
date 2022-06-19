package cmd

import (
	"github.com/GodYao1995/Goooooo/internal/admin/controller"
	"github.com/GodYao1995/Goooooo/pkg/casbin"
	"github.com/GodYao1995/Goooooo/pkg/db"
	"github.com/GodYao1995/Goooooo/pkg/xhttp"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func RunAdmin() {
	fx.New(injectAdmin()).Run()
}

func injectAdmin() fx.Option {
	return fx.Options(
		// Provide
		configModule,
		db.Module,
		xhttp.Module,
		casbin.Module,
		// Invoke
		controller.ModuleV1,
		controller.ModuleV2,
		// Options
		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)
}
