package v1

import (
	"context"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

const (
	USERINDEX = "user_index"
)

type UserEsRepo struct {
	index   string
	mapping string
	client  *elastic.Client
	log     *zap.Logger
}

func NewUserEs(client *elastic.Client, log *zap.Logger) domain.UserEsRepositoryFace {
	ue := &UserEsRepo{index: USERINDEX, mapping: "", client: client, log: log}
	err := _InitIndex(ue)
	if err != nil {
		ue.log.Warn(err.Error())
	}
	return ue
}

func (ue *UserEsRepo) CreateUserDocument(ctx context.Context, documents []*domain.UserEsPO) error {
	_, err := ue.client.Index().Index(ue.index).BodyJson(documents).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func _InitIndex(ue *UserEsRepo) error {
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
