package admin

import (
	"github.com/GodYao1995/Goooooo/internal/admin/controller"
	"github.com/GodYao1995/Goooooo/internal/config"
	"github.com/GodYao1995/Goooooo/pkg/casbin"
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
