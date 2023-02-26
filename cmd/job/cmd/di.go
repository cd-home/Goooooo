package cmd

import (
	"github.com/RichardKnop/machinery/v2"
	esrepo "github.com/cd-home/Goooooo/internal/admin/es"
	"github.com/cd-home/Goooooo/internal/admin/repository"
	job "github.com/cd-home/Goooooo/internal/job"
	tasks "github.com/cd-home/Goooooo/internal/job/tasks"
	"github.com/cd-home/Goooooo/pkg/db"
	"github.com/cd-home/Goooooo/pkg/logger"
	"github.com/cd-home/Goooooo/pkg/xes"
	xjob "github.com/cd-home/Goooooo/pkg/xjob"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var JobWorkerModule = fx.Invoke(JobWorker)

func Run() {
	fx.New(inject()).Run()
}

func JobWorker(jobServer *machinery.Server, _jobs *tasks.UserESTask) {
	jobServer.RegisterTask("sum", tasks.Sum)
	jobServer.RegisterTask("user2es", _jobs.UsersToES)
	worker := jobServer.NewWorker("worker", 0)
	worker.Launch()
}

func inject() fx.Option {
	return fx.Options(
		// Provide
		configModule,
		logger.Module,
		db.Module,
		xjob.Module,
		xes.Module,
		repository.Module,
		esrepo.Module,
		job.Module,
		// Invoke
		JobWorkerModule,
		// Options
		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)
}
