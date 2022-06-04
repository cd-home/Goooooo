package v1

import (
	"github.com/GodYao1995/Goooooo/internal/domain"
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

func (f FileLogic) UploadFile(fileName string, fileSize int64, fileUrl string, directory_id uint64, uploader uint64) error {
	return f.repo.UploadFile(fileName, fileSize, fileUrl, directory_id, uploader)
}

func (f FileLogic) DeleteFile(fileId uint64) error {
	return nil
}
