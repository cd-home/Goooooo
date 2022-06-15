package v1

import (
	"context"
	"time"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
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

func (repo RoleRepository) Create(ctx context.Context, roleId uint64, roleName string, roleLevel uint8, roleIndex uint8, father *uint64) (err error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "RoleRepository-Create")
	defer func() {
		span.SetTag("RoleRepository", "Create")
		span.Finish()
	}()
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

func (repo RoleRepository) Delete(ctx context.Context, roleId uint64) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "RoleRepository-Delete")
	defer func() {
		span.SetTag("RoleRepository", "Delete")
		span.Finish()
	}()
	var err error
	local := zap.Fields(zap.String("Repo", "DeleteRole"))
	_, err = repo.db.Exec(`UPDATE role SET delete_at = ? WHERE role_id = ?`, time.Now(), roleId)
	if err != nil {
		repo.log.WithOptions(local).Warn(err.Error())
		return err
	}
	// TODO: Should Delete RoleRelation ?
	return nil
}

func (repo RoleRepository) Update(ctx context.Context, roleId uint64, roleName string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "RoleRepository-Update")
	defer func() {
		span.SetTag("RoleRepository", "Update")
		span.Finish()
	}()
	var err error
	local := zap.Fields(zap.String("Repo", "DeleteRole"))
	_, err = repo.db.Exec(`UPDATE role SET role_name = ? WHERE role_id = ?`, roleName, roleId)
	if err != nil {
		repo.log.WithOptions(local).Warn(err.Error())
		return err
	}
	return nil
}

func (repo RoleRepository) Retrieve(ctx context.Context, roleLevel uint8, father *uint64) ([]*domain.RoleEntityDTO, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Retrieve")
	defer func() {
		span.SetTag("RoleRepository", "Retrieve")
		span.Finish()
	}()
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
		FROM 
		(
			SELECT
				role_id
			FROM role 
			WHERE role_id = ? AND delete_at is NULL
		) AS role
		JOIN
		(
			SELECT
				ancestor, descendant
			FROM role_relation 
			WHERE delete_at is NULL AND distance = 1
		) AS relation 
		ON role.role_id = relation.ancestor
		JOIN 
		(
			SELECT
				role_id, role_name, role_level, role_index, update_at, create_at
			FROM role 
			WHERE delete_at is NULL
		) AS son
		ON relation.descendant = son.role_id
		ORDER BY son.role_index`, *father)
	if err != nil {
		repo.log.WithOptions(local).Warn(err.Error())
		return nil, err
	}
	return roles, nil
}

func (repo RoleRepository) Move(ctx context.Context, roleId uint64, father uint64) (err error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "RoleRepository-Move")
	defer func() {
		span.SetTag("RoleRepository", "Move")
		span.Finish()
	}()
	var tx *sqlx.Tx
	local := zap.Fields(zap.String("Repo", "Move"))
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
	// Current Role_id`s All Childs
	relations := make([]*domain.RoleRelationPO, 0)
	err = tx.Select(&relations, `
		SELECT 
			ancestor, descendant, distance
		FROM role_relation WHERE ancestor = ? AND delete_at is NULL AND distance != 0`, roleId)
	if err != nil {
		repo.log.WithOptions(local).Warn(err.Error())
		return 
	}
	dests := make([]uint64, 0)
	for _, r := range relations {
		dests = append(dests, r.Descendant)
	}
	// 解除 roleId 与其 祖先 的关系
	tx.MustExec(`
		DELETE FROM role_relation 
		WHERE descendant = ? AND delete_at is NULL AND distance != 0`, roleId)

	// 解除 roleId 的孩子与 roleId 祖先的关系
	query, args, err := sqlx.In(`
		DELETE FROM role_relation 
		WHERE descendant in (?) AND delete_at is NULL AND distance > 1`, dests)

	if err != nil {
		return
	}
	tx.MustExec(query, args...)

	// Rebuild New Relation
	// 先创建于父级目录关系
	tx.MustExec(
		`INSERT INTO role_relation (ancestor, descendant, distance)
		 VALUES(?, ?, ?)`, father, roleId, 1)

	// 创建祖先与该目录的关系
	grandsRelations := make([]*domain.RoleRelationPO, 0)
	childsRelations := make([]*domain.RoleRelationPO, 0)

	tx.Select(&grandsRelations, `
		SELECT 
			ancestor, descendant, distance
		FROM role_relation WHERE descendant = ? AND distance != 0 AND delete_at is NULL`, father)

	// 移动到的father 本身就是顶层角色
	if len(grandsRelations) == 0 {
		for _, child := range relations {
			childsRelations = append(childsRelations, &domain.RoleRelationPO{
				Ancestor:   father,
				Descendant: child.Descendant,
				Distance:   child.Distance + 1,
			})
		}
		_, err = tx.NamedExec(`
			INSERT INTO role_relation (ancestor, descendant, distance) 
			VALUES(:ancestor, :descendant, :distance)`, childsRelations)
		if err != nil {
			repo.log.Sugar().Debug(childsRelations)
			repo.log.WithOptions(local).Warn(err.Error())
			return
		}
	}

	// 移动到的father 有 father
	if len(grandsRelations) > 0 {
		for _, relation := range grandsRelations {
			relation.Descendant = roleId
			relation.Distance = relation.Distance + 1
			for _, child := range relations {
				childsRelations = append(childsRelations, &domain.RoleRelationPO{
					Ancestor:   relation.Ancestor,
					Descendant: child.Descendant,
					Distance:   relation.Distance + child.Distance,
				})
			}
		}
		_, err = tx.NamedExec(`
			INSERT INTO role_relation (ancestor, descendant, distance) 
			VALUES(:ancestor, :descendant, :distance)`, grandsRelations)
		if err != nil {
			repo.log.Sugar().Debug(relations)
			repo.log.WithOptions(local).Warn(err.Error())
			return
		}
		_, err = tx.NamedExec(`
			INSERT INTO role_relation (ancestor, descendant, distance) 
			VALUES(:ancestor, :descendant, :distance)`, childsRelations)
		if err != nil {
			repo.log.Sugar().Debug(childsRelations)
			repo.log.WithOptions(local).Warn(err.Error())
			return 
		}
	}
	return 
}
