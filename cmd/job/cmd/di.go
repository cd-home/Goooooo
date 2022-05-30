package cmd

import (
	"github.com/GodYao1995/Goooooo/internal/job"
	jobs "github.com/GodYao1995/Goooooo/internal/job/jobs"
	"github.com/RichardKnop/machinery/v2"
	"go.uber.org/fx"
)

func Run() {
	fx.New(inject()).Run()
}

func JobWorker(jobServer *machinery.Server) {
	jobServer.RegisterTask("sum", jobs.Sum)
	worker := jobServer.NewWorker("worker", 0)
	worker.Launch()
}

var JobWorkerModule = fx.Invoke(JobWorker)

func inject() fx.Option {
	return fx.Options(
		// Provide
		job.Module,
		// Invoke
		JobWorkerModule,
		// Options
		// fx.WithLogger(
		// 	func() fxevent.Logger {
		// 		return fxevent.NopLogger
		// 	},
		// ),
	)
}
