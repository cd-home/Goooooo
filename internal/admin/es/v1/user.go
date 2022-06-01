package v1

import (
	"context"
	"errors"

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

// CreateUserDocument
func (ue *UserEsRepo) CreateUserDocument(ctx context.Context, document *domain.UserEsPO) error {
	_, err := ue.client.Index().Index(ue.index).BodyJson(document).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

// CreateUserDocuments Batch Add
func (ue *UserEsRepo) CreateUserDocuments(ctx context.Context, documents []*domain.UserEsPO) error {
	bulk := ue.client.Bulk().Index(ue.index)
	for _, doc := range documents {
		bulk.Add(elastic.NewBulkIndexRequest().Doc(doc))
		if bulk.NumberOfActions() >= len(documents) {
			// Commit
			res, err := bulk.Do(ctx)
			if err != nil {
				return err
			}
			if res.Errors {
				// Look up the failed documents with res.Failed(), and e.g. recommit
				return errors.New("bulk commit failed")
			}
			// "bulk" is reset after Do, so you can reuse it
		}
	}
	// Commit the final batch before exiting
	if bulk.NumberOfActions() > 0 {
		_, err := bulk.Do(ctx)
		if err != nil {
			return err
		}
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
