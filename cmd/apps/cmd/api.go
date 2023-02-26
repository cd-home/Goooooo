package cmd

import (
	"github.com/cd-home/Goooooo/internal/api/controller"
	"github.com/cd-home/Goooooo/pkg/casbin"
	"github.com/cd-home/Goooooo/pkg/db"
	"github.com/cd-home/Goooooo/pkg/xhttp"
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
		xhttp.Module,
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
