package job

import (
	tasks "github.com/GodYao1995/Goooooo/internal/job/tasks"
	"go.uber.org/fx"
)

var Module = fx.Provide(tasks.NewUserESJob)
