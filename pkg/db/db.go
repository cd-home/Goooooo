package db

import (
	"context"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewSqlx)

func NewSqlx(lifecycle fx.Lifecycle, vp *viper.Viper) *sqlx.DB {
	user, password := vp.GetString("DB.USER"), vp.GetString("DB.PASSWD")
	host, port, database := vp.GetString("DB.HOST"), vp.GetString("DB.PORT"), vp.GetString("DB.DATABASE")
	dns := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4"
	db, _ := sqlx.Open("mysql", dns)
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := db.Ping()
			if err != nil {
				log.Println(err)
			} else {
				db.SetMaxIdleConns(10)
				db.SetMaxOpenConns(128)
				db.SetConnMaxLifetime(time.Duration(7) * time.Hour)
				log.Printf("\033[1;32;32m=========== DB RUNNING: [ %s:%s:%s ] ============\033[0m", host, port, database)
			}
			return err
		},
		OnStop: func(ctx context.Context) error {
			log.Printf("\033[1;34;34m DB Close [ %s ] \033[0m\n", time.Now().Local())
			return db.Close()
		},
	})

	return db
}
