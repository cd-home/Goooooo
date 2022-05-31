package v1

import (
	"fmt"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"go.uber.org/zap"
)

type DirectoryrLogic struct {
	repo domain.DirectoryRepositoryFace
	log  *zap.Logger
}

func NewDirectoryrLogic(repo domain.DirectoryRepositoryFace, log *zap.Logger) domain.DirectoryLogicFace {
	return &DirectoryrLogic{repo: repo, log: log.WithOptions(zap.Fields(zap.String("module", "DirectoryrLogic")))}
}

// CreateDirectory
func (l *DirectoryrLogic) CreateDirectory(name string, dType string, level uint8, index uint8, father *uint64) error {
	return l.repo.CreateDirectory(name, dType, level, index, father)
}

// ListDirectory
func (l *DirectoryrLogic) ListDirectory(level uint8, father *uint64) []*domain.DirectoryVO {
	local := zap.Fields(zap.String("Logic", "ListDirectory"))
	objs := l.repo.ListDirectory(level, father)
	// 预估一个容量
	directoryVOs := make([]*domain.DirectoryVO, 0, 6)
	for _, obj := range objs {
		directoryVOs = append(directoryVOs, &domain.DirectoryVO{
			DirectoryId:    obj.DirectoryId,
			DirectoryName:  obj.DirectoryName,
			DirectoryType:  obj.DirectoryType,
			DirectoryLevel: obj.DirectoryLevel,
			DirectoryIndex: obj.DirectoryIndex,
		})
	}
	l.log.WithOptions(local).Debug("directoryVOs capacity", zap.String("directoryVOs caps", fmt.Sprint(cap(directoryVOs))))
	return directoryVOs
}

func (l *DirectoryrLogic) RenameDirectory(directory_id uint64, name string) *domain.DirectoryVO {
	local := zap.Fields(zap.String("Logic", "RenameDirectory"))
	obj := l.repo.RenameDirectory(directory_id, name)
	if obj != nil {
		directoryVO := &domain.DirectoryVO{
			DirectoryId:    obj.DirectoryId,
			DirectoryName:  obj.DirectoryName,
			DirectoryType:  obj.DirectoryType,
			DirectoryLevel: obj.DirectoryLevel,
			DirectoryIndex: obj.DirectoryIndex,
		}
		return directoryVO
	}
	l.log.WithOptions(local).Debug("directoryVOs Dont exist", zap.String("directory_id", fmt.Sprint(directory_id)))
	return nil
}
