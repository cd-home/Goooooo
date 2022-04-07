package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/fx"
)

var Module = fx.Provide(New)

func New(lifecycle fx.Lifecycle, vp *viper.Viper) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())

	engine.GET("/docs/*any", swagger.WrapHandler(swaggerFiles.Handler))
	app := vp.GetString("APP")

	srv := &http.Server{
		Addr:         vp.GetString(app + ".SERVER_HOST"),
		Handler:      engine,
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
	}
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil {
					log.Println(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping HTTP server.")
			return srv.Shutdown(ctx)
		},
	})
	return engine
}
