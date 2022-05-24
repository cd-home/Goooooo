package xes

import (
	"context"
	"log"

	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewESClient)

func NewESClient(lifecycle fx.Lifecycle, vp *viper.Viper) *elastic.Client {
	addr, user, pwd := vp.GetString("ES.ADDR"), vp.GetString("ES.USER"), vp.GetString("ES.PASSWD")
	log.Println(addr, user, pwd)
	var client *elastic.Client
	var err error
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			client, err = elastic.NewClient(
				elastic.SetURL(addr),
				elastic.SetSniff(false),
				// elastic.SetBasicAuth(user, pwd),
			)
			if err != nil {
				log.Println(err)
				return err
			}
			_, _, err = client.Ping(addr).Do(context.Background())
			if err != nil {
				log.Println(err)
				return err
			}
			v, _ := client.ElasticsearchVersion(addr)
			log.Printf("\033[1;32;32m=========== ES [%s] RUNNING: [ %s ] ============\033[0m", v, addr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			client.Stop()
			return nil
		},
	})
	log.Println(client)
	return client
}
