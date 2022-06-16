package v1

import (
	"context"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/pkg/tools"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type FileRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewFileRepository(db *sqlx.DB, log *zap.Logger) domain.FileRepositoryFace {
	return &FileRepository{
		db:  db,
		log: log.WithOptions(zap.Fields(zap.String("module", "FileRepository"))),
	}
}

func (f FileRepository) UploadFile(ctx context.Context, fileName string, fileSize int64, fileUrl string, directoryId uint64, uploader uint64) error {
	var err error
	local := zap.Fields(zap.String("Repo", "UploadFile"))
	_, err = f.db.Exec(`
		INSERT INTO file (file_id, file_name, file_size, file_url, directory_id, uploader) 
		VALUES(?, ?, ?, ?, ?, ?)`, tools.SnowId(), fileName, fileSize, fileUrl, directoryId, uploader)

	if err != nil {
		f.log.WithOptions(local).Warn(err.Error())
		return err
	}
	return nil
}

func (f FileRepository) DeleteFile(fileId uint64) error {
	return nil
}
