package v1

import (
	"context"
	"time"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/pkg/tools"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type DirectoryRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewDirectoryRepository(db *sqlx.DB, log *zap.Logger) domain.DirectoryRepositoryFace {
	return &DirectoryRepository{db: db, log: log.WithOptions(zap.Fields(zap.String("module", "DirectoryRepository")))}
}

func (repo DirectoryRepository) Create(ctx context.Context, name string, dType string, level uint8, index uint8, father *uint64) (err error) {
	var tx *sqlx.Tx
	local := zap.Fields(zap.String("Repo", "CreateDirectory"))
	// 是否考虑换一种唯一资源标识
	unique := uint64(tools.SnowId())
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
	// Directory
	tx, err = repo.db.Beginx()
	tx.MustExec(
		`INSERT INTO directory(directory_id, directory_name, directory_type, directory_level, directory_index) 
		 VALUES(?, ?, ?, ?, ?)`, unique, name, dType, level, index)

	// DirectoryRelation
	tx.MustExec(
		`INSERT INTO directory_relation(ancestor, descendant, distance)
		 VALUES(?, ?, ?)`, unique, unique, 0)

	// 创建下级目录
	if father != nil {

		// 先创建于父级目录关系
		tx.MustExec(
			`INSERT INTO directory_relation(ancestor, descendant, distance)
			 VALUES(?, ?, ?)`, *father, unique, 1)

		// 创建祖先与该目录的关系
		relations := make([]*domain.DirectoryRelationPO, 0)
		tx.Select(&relations, `
			SELECT 
				ancestor, descendant, distance
			FROM directory_relation WHERE descendant = ? AND distance != 0`, *father)
		if len(relations) > 0 {
			for _, relation := range relations {
				relation.Descendant = unique
				relation.Distance = relation.Distance + 1
			}
			_, err = tx.NamedExec(
				`INSERT INTO directory_relation(ancestor, descendant, distance) 
				VALUES(:ancestor, :descendant, :distance)`, relations)
		}
	}
	return err
}

func (repo DirectoryRepository) Delete(ctx context.Context, directory_id uint64) error {
	var err error
	local := zap.Fields(zap.String("Repo", "Delete"))
	_, err = repo.db.Exec(`UPDATE directory SET delete_at = ? WHERE directory_id = ?`, time.Now(), directory_id)
	if err != nil {
		repo.log.WithOptions(local).Warn(err.Error())
		return err
	}
	// TODO: Should Delete RoleRelation ?
	return nil
}

func (repo DirectoryRepository) Update(ctx context.Context, directory_id uint64, name string) *domain.DirectoryDTO {
	var err error
	local := zap.Fields(zap.String("Repo", "Update"))
	_, err = repo.db.Exec(`UPDATE directory SET directory_name = ? WHERE directory_id = ?`, name, directory_id)
	if err != nil {
		repo.log.WithOptions(local).Warn(err.Error())
		return nil
	}
	// An error is returned if the result set is empty.
	var directory domain.DirectoryDTO
	err = repo.db.Get(&directory, `
		SELECT 
			directory_id, directory_name, directory_type, directory_level, directory_index 
		FROM directory WHERE directory_id = ?`, directory_id)
	if err != nil {
		return nil
	}
	return &directory
}

func (repo DirectoryRepository) Retrieve(ctx context.Context, level uint8, father *uint64) []*domain.DirectoryDTO {
	// First Class Directory
	directories := make([]*domain.DirectoryDTO, 0)
	if father == nil && level == 1 {
		repo.db.Select(&directories, `
			SELECT 
				directory_id, directory_name, directory_type, directory_level, directory_index
			FROM directory WHERE directory_level = ? AND delete_at is NULL ORDER BY directory_index`, level)
		return directories
	}
	// Other Way Do not Need Level
	//  SELECT
	//		d2.directory_id, d2.directory_name, d2.directory_type, d2.directory_level, d2.directory_index
	//	FROM directory AS d
	//	JOIN directory_relation AS relation ON d.directory_id = relation.ancestor
	//	JOIN directory AS d2 ON relation.descendant = d2.directory_id
	//	WHERE d.directory_id = ? AND d.delete_at IS NULL AND relation.distance = 1
	repo.db.Select(&directories, `
		SELECT 
			directory_id, directory_name, directory_type, directory_level, directory_index
		FROM directory as directory JOIN
		(
			SELECT
				descendant
			FROM directory_relation
			WHERE ancestor = ? AND delete_at is NULL
		) as relation
		ON directory.directory_id = relation.descendant
		WHERE directory.directory_level = ? AND directory.delete_at is NULL
		ORDER BY directory.directory_index`, *father, level)
	return directories
}

func (repo DirectoryRepository) Move(ctx context.Context, directory_id uint64, father uint64) (err error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "DirectoryRepository-Move")
	defer func() {
		span.SetTag("DirectoryrLogic", "MoveDirectory")
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
		return
	}
	relations := make([]*domain.DirectoryRelationPO, 0)
	err = tx.Select(&relations, `
		SELECT 
			ancestor, descendant, distance
		FROM directory WHERE delete_at is NULL AND ancestor = ?`, directory_id)
	if err != nil {
		repo.log.WithOptions(local).Warn(err.Error())
		return
	}

	// Current directory_id with ancestors => break up
	tx.MustExec(`
		DELETE FROM directory_relation 
		WHERE descendant = ? AND delete_at is NULL AND distance != 0`, directory_id)

	// Current directory_id`s Childs with directory_id`s ancestors => break up
	var dests []uint64
	for _, r := range relations {
		dests = append(dests, r.Descendant)
	}
	query, args, err := sqlx.In(`
		DELETE FROM directory_relation 
		WHERE descendant in (?) AND delete_at is NULL AND distance > 1`, dests)
	if err != nil {
		return
	}
	tx.MustExec(query, args...)

	// Rebuild Relation
	tx.MustExec(
		`INSERT INTO directory_relation (ancestor, descendant, distance)
		 VALUES(?, ?, ?)`, father, directory_id, 1)

	grandsRelations := make([]*domain.DirectoryRelationPO, 0)
	childsRelations := make([]*domain.DirectoryRelationPO, 0)

	tx.Select(&grandsRelations, `
		SELECT 
			ancestor, descendant, distance
		FROM directory_relation WHERE descendant = ? AND distance != 0 AND delete_at is NULL`, father)

	// 移动到的目标father 本身就是顶层目录
	if len(grandsRelations) == 0 {
		for _, child := range relations {
			childsRelations = append(childsRelations, &domain.DirectoryRelationPO{
				Ancestor:   father,
				Descendant: child.Descendant,
				Distance:   child.Distance + 1,
			})
		}
		_, err = tx.NamedExec(`
			INSERT INTO directory_relation (ancestor, descendant, distance) 
			VALUES(:ancestor, :descendant, :distance)`, childsRelations)
		if err != nil {
			repo.log.Sugar().Debug(childsRelations)
			repo.log.WithOptions(local).Warn(err.Error())
			return
		}
		return
	}

	// 移动到的father 有 father
	for _, relation := range grandsRelations {
		relation.Descendant = directory_id
		relation.Distance = relation.Distance + 1
		for _, child := range relations {
			childsRelations = append(childsRelations, &domain.DirectoryRelationPO{
				Ancestor:   relation.Ancestor,
				Descendant: child.Descendant,
				Distance:   relation.Distance + child.Distance,
			})
		}
	}
	// father => childs 需要单独添加
	for _, child := range relations {
		childsRelations = append(childsRelations, &domain.DirectoryRelationPO{
			Ancestor:   father,
			Descendant: child.Descendant,
			Distance:   child.Distance + 1,
		})
	}
	_, err = tx.NamedExec(`
			INSERT INTO directory_relation (ancestor, descendant, distance) 
			VALUES(:ancestor, :descendant, :distance)`, grandsRelations)
	if err != nil {
		repo.log.Sugar().Debug(relations)
		repo.log.WithOptions(local).Warn(err.Error())
		return
	}
	_, err = tx.NamedExec(`
			INSERT INTO directory_relation (ancestor, descendant, distance) 
			VALUES(:ancestor, :descendant, :distance)`, childsRelations)
	if err != nil {
		repo.log.Sugar().Debug(childsRelations)
		repo.log.WithOptions(local).Warn(err.Error())
		return
	}
	return
}
