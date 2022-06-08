package v1

import (
	"context"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/pkg/tools"
	"go.uber.org/zap"
)

type RoleLogic struct {
	repo domain.RoleRepositoryFace
	log  *zap.Logger
}

func NewRoleLogic(repo domain.RoleRepositoryFace, log *zap.Logger) domain.RoleLogicFace {
	return &RoleLogic{repo: repo, log: log}
}

func (r RoleLogic) CreateRole(ctx context.Context, roleName string, roleLevel uint8, roleIndex uint8, parent *uint64) error {
	return r.repo.CreateRole(ctx, uint64(tools.SnowId()), roleName, roleLevel, roleIndex, parent)
}
