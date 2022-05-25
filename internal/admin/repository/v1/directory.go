package v1

import (
	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/pkg/tools"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type DirectoryRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewDirectoryRepository(db *sqlx.DB, log *zap.Logger) domain.DirectoryRepositoryFace {
	return &DirectoryRepository{db: db, log: log.WithOptions(zap.Fields(zap.String("module", "DirectoryRepository")))}
}

// CreateDirectory
func (repo *DirectoryRepository) CreateDirectory(name string, dType string, level uint8, index uint8, father *uint64) (err error) {
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

// ListDirectory GET Direct Children Directory
func (repo *DirectoryRepository) ListDirectory(level uint8, father *uint64) []*domain.DirectoryDTO {
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
	//	JOINWHERE d.directory_id = ? AND d.delete_at IS NULL AND relation.distance = 1
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

func (repo *DirectoryRepository) RenameDirectory(directory_id uint64, name string) *domain.DirectoryDTO {
	var err error
	local := zap.Fields(zap.String("Repo", "RenameDirectory"))
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
