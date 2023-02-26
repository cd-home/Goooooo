package job

import (
	tasks "github.com/cd-home/Goooooo/internal/job/tasks"
	"go.uber.org/fx"
)

var Module = fx.Provide(tasks.NewUserESJob)
