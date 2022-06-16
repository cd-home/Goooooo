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

// UploadFile
func (f FileLogic) UploadFile(ctx context.Context, fileName string, fileSize int64, fileType string, fileUrl string, directory_id uint64, uploader uint64) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "FileLogic-UploadFile")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("FileLogic", "UploadFile")
		span.Finish()
	}()
	return f.repo.UploadFile(next, fileName, fileSize, fileType, fileUrl, directory_id, uploader)
}

// DeleteFile
func (f FileLogic) DeleteFile(ctx context.Context, fileId uint64) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "FileLogic-DeleteFile")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("FileLogic", "DeleteFile")
		span.Finish()
	}()
	return f.repo.DeleteFile(next, fileId)
}

// RetrieveFiles
func (f FileLogic) RetrieveFiles(ctx context.Context) ([]*domain.FileEntityVo, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "FileController-ListFile")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("FileController", "ListFile")
		span.Finish()
	}()
	dtos, err := f.repo.RetrieveFiles(next)
	if err != nil {
		return nil, err
	}
	vos := make([]*domain.FileEntityVo, 0)
	for _, obj := range dtos {
		vos = append(vos, &domain.FileEntityVo{
			FileId:    obj.FileId,
			FileName:  obj.FileName,
			FileSize:  obj.FileSize,
			FileType:  obj.FileType,
			FileUrl:   obj.FileUrl,
			Directory: obj.Directory,
			CreateAt:  obj.CreateAt,
			UpdateAt:  obj.UpdateAt,
			Uploader:  obj.Uploader,
		})
	}
	return vos, nil
}
