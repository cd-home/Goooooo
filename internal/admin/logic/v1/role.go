package v1

import (
	"context"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"go.uber.org/zap"
)

type RoleLogic struct {
	repo domain.RoleRepositoryFace
	log  *zap.Logger
}

func NewRoleLogic(repo domain.RoleRepositoryFace, log *zap.Logger) domain.RoleLogicFace {
	return &RoleLogic{repo: repo, log: log}
}

func (r RoleLogic) CreateRole(ctx context.Context) {

}
