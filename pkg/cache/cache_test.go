package cache

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/GodYao1995/Goooooo/pkg/config"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func TestNewRedisClient(t *testing.T) {
	var newVp func() *viper.Viper = func() *viper.Viper {
		return config.NewViper("admin", "dev", "../config/testdata/configs")
	}
	app := fx.New(
		fx.Provide(
			newVp,
		),
		fx.Invoke(NewRedisClient),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
	stopCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
