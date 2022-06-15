package v1

import (
	"context"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/pkg/tools"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type RoleLogic struct {
	repo domain.RoleRepositoryFace
	log  *zap.Logger
}

func NewRoleLogic(repo domain.RoleRepositoryFace, log *zap.Logger) domain.RoleLogicFace {
	return &RoleLogic{repo: repo, log: log}
}

func (r RoleLogic) CreateRole(ctx context.Context, roleName string, roleLevel uint8, roleIndex uint8, father *uint64) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "RoleController-CreateRole")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("RoleLogic", "CreateRole")
		span.Finish()
	}()
	return r.repo.Create(next, uint64(tools.SnowId()), roleName, roleLevel, roleIndex, father)
}

func (r RoleLogic) RetrieveRoles(ctx context.Context, roleLevel uint8, father *uint64) ([]*domain.RoleEntityVO, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "RetrieveRoles")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("RoleLogic", "RetrieveRoles")
		span.Finish()
	}()
	dto, err := r.repo.Retrieve(next, roleLevel, father)
	if err != nil {
		return nil, err
	}
	roleVos := make([]*domain.RoleEntityVO, 0)
	for _, obj := range dto {
		roleVos = append(roleVos, &domain.RoleEntityVO{
			RoleId:    obj.RoleId,
			RoleName:  obj.RoleName,
			RoleLevel: obj.RoleLevel,
			RoleIndex: obj.RoleIndex,
			CreateAt:  obj.CreateAt,
			UpdateAt:  obj.UpdateAt,
		})
	}
	return roleVos, nil
}

func (r RoleLogic) DeleteRole(ctx context.Context, roleId uint64) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "RoleLogic-DeleteRole")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("RoleLogic", "DeleteRole")
		span.Finish()
	}()
	return r.repo.Delete(next, roleId)
}

func (r RoleLogic) UpdateRole(ctx context.Context, roleId uint64, roleName string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "RoleLogic-UpdateRole")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("RoleLogic", "UpdateRole")
		span.Finish()
	}()
	return r.repo.Update(next, roleId, roleName)
}

func (r RoleLogic) MoveRole(ctx context.Context, roleId uint64, father uint64) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "RoleLogic-MoveRole")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("RoleLogic", "MoveRole")
		span.Finish()
	}()
	return r.repo.Move(next, roleId, father)
}
