package v1

import (
	"context"

	"github.com/cd-home/Goooooo/internal/domain"
	"github.com/cd-home/Goooooo/pkg/tools"
	"github.com/cd-home/Goooooo/pkg/xtime"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
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

// UploadFile
func (f FileRepository) UploadFile(ctx context.Context, fileName string, fileSize int64, fileType string, fileUrl string, directoryId uint64, uploader uint64) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "FileRepository-UploadFile")
	defer func() {
		span.SetTag("FileRepository", "UploadFile")
		span.Finish()
	}()
	var err error
	local := zap.Fields(zap.String("Repo", "UploadFile"))
	_, err = f.db.Exec(`
		INSERT INTO file (file_id, file_name, file_size, file_type, file_url, directory_id, uploader) 
		VALUES(?, ?, ?, ?, ?, ?, ?)`, tools.SnowId(), fileName, fileSize, fileType, fileUrl, directoryId, uploader)

	if err != nil {
		f.log.WithOptions(local).Warn(err.Error())
		return err
	}
	return nil
}

// DeleteFile
func (f FileRepository) DeleteFile(ctx context.Context, fileId uint64) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "FileRepository-DeleteFile")
	defer func() {
		span.SetTag("FileRepository", "DeleteFile")
		span.Finish()
	}()
	var err error
	local := zap.Fields(zap.String("Repo", "DeleteFile"))
	_, err = f.db.Exec(`UPDATE file SET delete_at = ? WHERE file_id = ?`, xtime.Now(), fileId)
	if err != nil {
		f.log.WithOptions(local).Warn(err.Error())
		return err
	}
	return nil
}

// Retrieve
func (f FileRepository) RetrieveFiles(ctx context.Context) ([]*domain.FileEntityDTO, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "FileRepository-RetrieveFiles")
	defer func() {
		span.SetTag("FileRepository", "RetrieveFiles")
		span.Finish()
	}()
	var err error
	local := zap.Fields(zap.String("Repo", "RetrieveFiles"))
	files := make([]*domain.FileEntityDTO, 0)
	err = f.db.Select(&files, `
		SELECT 
			file_id, file_name, file_size, file_type, uploader, directory_id, file_url, update_at, create_at
		FROM file WHERE delete_at is NULL`)
	if err != nil {
		f.log.WithOptions(local).Warn(err.Error())
		return nil, err
	}
	return files, nil
}

// Retrieve
func (f FileRepository) RetrieveFilesByFather(ctx context.Context) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "FileRepository-RetrieveFilesByFather")
	defer func() {
		span.SetTag("FileRepository", "RetrieveFilesByFather")
		span.Finish()
	}()
	return nil
}
