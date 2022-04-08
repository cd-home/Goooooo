package db

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/GodYao1995/Goooooo/pkg/config"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func TestNewSqlxDB(t *testing.T) {
	var newVp func() *viper.Viper = func() *viper.Viper {
		return config.NewViper("admin", "dev", "../config/testdata/configs")
	}
	app := fx.New(
		fx.Provide(
			newVp,
		),
		fx.Invoke(NewSqlx),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
	stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}

func TestProdNewSqlxDB(t *testing.T) {
	godotenv.Load("../config/testdata/.env")
	var newEnvVp func() *viper.Viper = func() *viper.Viper {
		return config.NewViper(os.Getenv("APP_NAME"), os.Getenv("APP_MODE"))
	}
	app := fx.New(
		fx.Provide(
			newEnvVp,
		),
		fx.Invoke(NewSqlx),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
	stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
