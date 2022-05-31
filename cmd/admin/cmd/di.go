package cmd

import (
	"github.com/GodYao1995/Goooooo/internal/admin/controller"
	esrepo "github.com/GodYao1995/Goooooo/internal/admin/es"
	"github.com/GodYao1995/Goooooo/internal/admin/logic"
	"github.com/GodYao1995/Goooooo/internal/admin/repository"
	"github.com/GodYao1995/Goooooo/internal/admin/version"
	"github.com/GodYao1995/Goooooo/internal/pkg/session"
	"github.com/GodYao1995/Goooooo/pkg/cache"
	"github.com/GodYao1995/Goooooo/pkg/casbin"
	"github.com/GodYao1995/Goooooo/pkg/db"
	"github.com/GodYao1995/Goooooo/pkg/logger"
	"github.com/GodYao1995/Goooooo/pkg/xes"
	"github.com/GodYao1995/Goooooo/pkg/xhttp/server"
	xjob "github.com/GodYao1995/Goooooo/pkg/xjob"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func Run() {
	fx.New(inject()).Run()
}

func inject() fx.Option {
	return fx.Options(
		// Provide
		configModule,
		logger.Module,
		db.Module,
		cache.Module,
		session.Module,
		server.Module,
		version.Module,
		casbin.Module,
		xes.Module,
		xjob.Module,
		// Invoke
		controller.ModuleV1,
		controller.ModuleV2,
		// Provide
		logic.Module,
		repository.Module,
		esrepo.Module,
		// Options
		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)
}
