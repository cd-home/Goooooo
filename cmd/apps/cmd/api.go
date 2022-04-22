package cmd

import (
	"github.com/GodYao1995/Goooooo/internal/api/controller"
	"github.com/GodYao1995/Goooooo/pkg/casbin"
	"github.com/GodYao1995/Goooooo/pkg/db"
	"github.com/GodYao1995/Goooooo/pkg/xhttp/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func RunApi() {
	fx.New(injectApi()).Run()
}

func injectApi() fx.Option {
	return fx.Options(
		// Provide
		configModule,
		db.Module,
		server.Module,
		casbin.Module,
		// Invoke
		controller.Module,
		// Options
		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)
}
