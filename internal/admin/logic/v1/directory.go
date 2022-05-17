package v1

import "github.com/GodYao1995/Goooooo/internal/domain"

type DirectoryrLogic struct {
	repo domain.DirectoryRepositoryFace
}

func NewDirectoryrLogic(repo domain.DirectoryRepositoryFace) domain.DirectoryLogicFace {
	return &DirectoryrLogic{repo: repo}
}

func (l *DirectoryrLogic) CreateDirectory(name string, dType string, level uint8, index uint8, father *uint64) error {
	return l.repo.CreateDirectory(name, dType, level, index, father)
}

func (l *DirectoryrLogic) ListDirectory(level uint8, father *uint64) []*domain.DirectoryVO {
	objs := l.repo.ListDirectory(level, father)
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
	return directoryVOs
}
