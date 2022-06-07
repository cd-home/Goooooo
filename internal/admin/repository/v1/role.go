package v1

import (
	"context"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type RoleRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewRoleRepository(db *sqlx.DB, log *zap.Logger) domain.RoleRepositoryFace {
	return &RoleRepository{
		db:  db,
		log: log.WithOptions(zap.Fields(zap.String("module", "RoleRepository"))),
	}
}

func (repo RoleRepository) CreateRole(ctx context.Context) error {
	return nil
}

func (repo RoleRepository) CreateRelation(ctx context.Context) error {
	return nil
}

func (repo RoleRepository) Delete(ctx context.Context) {
}

func (repo RoleRepository) Update(ctx context.Context) {
}

func (repo RoleRepository) Retrieve(ctx context.Context) {

}
