package v1

import (
	"context"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type FileLogic struct {
	repo domain.FileRepositoryFace
	log  *zap.Logger
}

func NewFileLogic(repo domain.FileRepositoryFace, log *zap.Logger) domain.FileLogicFace {
	return &FileLogic{
		repo: repo,
		log:  log.WithOptions(zap.Fields(zap.String("module", "FileLogic"))),
	}
}

func (f FileLogic) UploadFile(ctx context.Context, fileName string, fileSize int64, fileUrl string, directory_id uint64, uploader uint64) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "FileLogic-UploadFile")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("FileLogic", "UploadFile")
		span.Finish()
	}()
	return f.repo.UploadFile(next, fileName, fileSize, fileUrl, directory_id, uploader)
}

func (f FileLogic) DeleteFile(fileId uint64) error {
	return nil
}
