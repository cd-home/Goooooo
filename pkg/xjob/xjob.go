package job

import (
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewJobServer)

func NewJobServer(vp *viper.Viper) *machinery.Server {
	cnf := &config.Config{
		DefaultQueue:    vp.GetString("JOB.QUEUE"),
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
	broker := redisbroker.NewGR(cnf, vp.GetStringSlice("JOB.BROKER"), vp.GetInt("JOB.BROKERDB"))
	backend := redisbackend.NewGR(cnf, vp.GetStringSlice("JOB.BACKEND"), vp.GetInt("JOB.BACKENDDB"))
	lock := eagerlock.New()
	JobServer := machinery.NewServer(cnf, broker, backend, lock)
	return JobServer
}
