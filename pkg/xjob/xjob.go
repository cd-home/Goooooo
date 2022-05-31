package job

import (
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewJobServer)

func NewJobServer() *machinery.Server {
	cnf := &config.Config{
		Broker:          "",
		DefaultQueue:    "machinery_tasks",
		ResultsExpireIn: 3600,
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  3000,
			DelayedTasksPollPeriod: 500,
		},
	}
	broker := redisbroker.NewGR(cnf, []string{"127.0.0.1:6379"}, 0)
	backend := redisbackend.NewGR(cnf, []string{"127.0.0.1:6379"}, 0)
	lock := eagerlock.New()
	JobServer := machinery.NewServer(cnf, broker, backend, lock)
	return JobServer
}
