package cmd

import (
	"github.com/cd-home/Goooooo/internal/admin/controller"
	esrepo "github.com/cd-home/Goooooo/internal/admin/es"
	"github.com/cd-home/Goooooo/internal/admin/logic"
	"github.com/cd-home/Goooooo/internal/admin/repository"
	"github.com/cd-home/Goooooo/internal/admin/version"
	"github.com/cd-home/Goooooo/internal/pkg/session"
	"github.com/cd-home/Goooooo/pkg/cache"
	"github.com/cd-home/Goooooo/pkg/casbin"
	"github.com/cd-home/Goooooo/pkg/db"
	"github.com/cd-home/Goooooo/pkg/logger"
	"github.com/cd-home/Goooooo/pkg/xes"
	"github.com/cd-home/Goooooo/pkg/xhttp"
	xjob "github.com/cd-home/Goooooo/pkg/xjob"
	"github.com/cd-home/Goooooo/pkg/xtracer"
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
		xhttp.Module,
		version.Module,
		casbin.Module,
		xes.Module,
		xjob.Module,
		xtracer.Module,
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
