package v1

import (
	"context"

	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

type UserEsRepo struct {
	index   string
	mapping string
	client  *elastic.Client
	log     *zap.Logger
}

func NewUserEs(index string, client *elastic.Client, log *zap.Logger) *UserEsRepo {
	ue := &UserEsRepo{index: index, mapping: "", client: client, log: log}
	err := ue.InitIndex()
	if err != nil {
		ue.log.Warn(err.Error())
	}
	return ue
}

func (ue *UserEsRepo) InitIndex() error {
	ctx := context.Background()
	exist, err := ue.client.IndexExists(ue.index).Do(ctx)
	if err != nil {
		return err
	}
	if !exist {
		if len(ue.mapping) > 0 {
			_, err = ue.client.CreateIndex(ue.index).BodyString(ue.mapping).Do(ctx)
			return err
		}
		_, err = ue.client.CreateIndex(ue.index).Do(ctx)
		return err
	}
	return nil
}
