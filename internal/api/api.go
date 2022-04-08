package api

import (
	"github.com/GodYao1995/Goooooo/internal/api/controller"
	"github.com/GodYao1995/Goooooo/internal/config"
	"github.com/GodYao1995/Goooooo/pkg/db"
	"github.com/GodYao1995/Goooooo/pkg/xhttp/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func Run() {
	fx.New(inject()).Run()
}

func inject() fx.Option {
	return fx.Options(
		// Provide
		config.Module,
		db.Module,
		server.Module,

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
