package job

import (
	jobs "github.com/GodYao1995/Goooooo/internal/job/jobs"
	"go.uber.org/fx"
)

var Module = fx.Provide(jobs.NewUserESJob)
