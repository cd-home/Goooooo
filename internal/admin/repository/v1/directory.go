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

func (repo *DirectoryRepository) CreateDirectory(name string, dType string, level uint8, index uint8, father *uint64) error {
	var err error
	var tx *sqlx.Tx
	local := zap.Fields(zap.String("Repo", "CreateDirectory"))
	unique := uint64(tools.Ids())
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			repo.log.WithOptions(local).Warn(err.Error())
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

	if father != nil {

		tx.MustExec(
			`INSERT INTO directory_relation(ancestor, descendant, distance)
			 VALUES(?, ?, ?)`, *father, unique, 1)

		relations := make([]*domain.DirectoryRelation, 0)
		tx.Select(&relations, `
			SELECT 
				ancestor,descendant,distance
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
