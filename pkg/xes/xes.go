package xes

import (
	"context"
	"log"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewESClient)

func NewESClient(lifecycle fx.Lifecycle, vp *viper.Viper) *elastic.Client {
	addr, _, _ := vp.GetString("ES.ADDR"), vp.GetString("ES.USER"), vp.GetString("ES.PASSWD")
	var err error
	client, err := elastic.NewClient(
		elastic.SetURL(addr),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(60*time.Second),
		// elastic.SetBasicAuth(user, pwd),
	)
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			_, _, err = client.Ping(addr).Do(context.Background())
			if err != nil {
				return err
			}
			v, _ := client.ElasticsearchVersion(addr)
			log.Printf("\033[1;32;32m=========== ES [%s] Running: [ %s ] \033[0m", v, addr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			client.Stop()
			return nil
		},
	})
	return client
}
