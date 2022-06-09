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

func (repo RoleRepository) CreateRole(ctx context.Context, roleId uint64, roleName string, roleLevel uint8, roleIndex uint8, father *uint64) (err error) {
	var tx *sqlx.Tx
	local := zap.Fields(zap.String("Repo", "CreateRole"))
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			repo.log.WithOptions(local).Error(err.Error())
			tx.Rollback()
		} else if err != nil {
			repo.log.WithOptions(local).Warn(err.Error())
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	tx, err = repo.db.Beginx()
	if err != nil {
		repo.log.WithOptions(local).Warn(err.Error())
		return
	}
	tx.MustExec(
		`INSERT INTO role (role_id, role_name, role_level, role_index) 
		 VALUES(?, ?, ?, ?)`, roleId, roleName, roleLevel, roleIndex)

	tx.MustExec(
		`INSERT INTO role_relation (ancestor, descendant, distance)
		 VALUES(?, ?, ?)`, roleId, roleId, 0)

	if father != nil {
		// 先创建于父级目录关系
		tx.MustExec(
			`INSERT INTO role_relation (ancestor, descendant, distance)
			 VALUES(?, ?, ?)`, *father, roleId, 1)
		// 创建祖先与该目录的关系
		relations := make([]*domain.RoleRelationPO, 0)
		tx.Select(&relations, `
			SELECT 
				ancestor, descendant, distance
			FROM role_relation WHERE descendant = ? AND distance != 0`, *father)
		if len(relations) > 0 {
			for _, relation := range relations {
				relation.Descendant = roleId
				relation.Distance = relation.Distance + 1
			}
			_, err = tx.NamedExec(
				`INSERT INTO role_relation (ancestor, descendant, distance) 
				VALUES(:ancestor, :descendant, :distance)`, relations)
			if err != nil {
				repo.log.WithOptions(local).Warn(err.Error())
				return
			}
		}
	}
	return
}

func (repo RoleRepository) CreateRelation(ctx context.Context) error {
	return nil
}

func (repo RoleRepository) Delete(ctx context.Context) {
}

func (repo RoleRepository) Update(ctx context.Context) {
}

func (repo RoleRepository) Retrieve(ctx context.Context, roleLevel uint8, father *uint64) ([]*domain.RoleEntityDTO, error) {
	var err error
	local := zap.Fields(zap.String("Repo", "CreateRole"))
	roles := make([]*domain.RoleEntityDTO, 0)
	// 一级角色
	if father == nil && roleLevel == 1 {
		err = repo.db.Select(&roles, `
			SELECT	
				role_id, role_name, role_level, role_index, role.update_at, role.create_at
			FROM role WHERE role_level = ? AND delete_at is NULL ORDER BY role_index`, roleLevel)
		if err != nil {
			repo.log.WithOptions(local).Warn(err.Error())
			return nil, err
		}
		return roles, nil
	}
	// 子角色
	err = repo.db.Select(&roles, `
		SELECT
			son.role_id, son.role_name, son.role_level, son.role_index, son.update_at, son.create_at
		FROM role as role
		JOIN role_relation as relation ON role.role_id = relation.ancestor
		JOIN role as son ON relation.descendant = son.role_id
		WHERE role.role_id = ? 
		AND role.delete_at is NULL 
		AND relation.delete_at is NULL
		AND son.delete_at is NULL 
		AND relation.distance = 1
		ORDER BY son.role_index`, *father)
	if err != nil {
		repo.log.WithOptions(local).Warn(err.Error())
		return nil, err
	}
	return roles, nil
}
