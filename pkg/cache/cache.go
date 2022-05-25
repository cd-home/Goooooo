package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewRedisClient)

func NewRedisClient(lifecycle fx.Lifecycle, vp *viper.Viper) *redis.Client {
	opts := &redis.Options{
		Addr:     vp.GetString("REDIS.ADDR"),
		Password: vp.GetString("REDIS.PASSWD"), // no password set
		DB:       vp.GetInt("REDIS.DB"),        // use default DB
	}
	rdb := redis.NewClient(opts)
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := rdb.Ping(context.Background()).Err(); err != nil {
				log.Fatal(err)
				return err
			}
			log.Printf("\033[1;32;32m=========== Redis  Running: [ %s ] \033[0m", opts.Addr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return rdb.Close()
		},
	})
	return rdb
}
