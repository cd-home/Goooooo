package domain

import "context"

type DirectoryLogicFace interface {
	CreateDirectory(ctx context.Context, name string, dType string, level uint8, index uint8, father *uint64) error
	ListDirectory(ctx context.Context, level uint8, directory_id *uint64) []*DirectoryVO
	RenameDirectory(ctx context.Context, directory_id uint64, name string) *DirectoryVO
}

type DirectoryRepositoryFace interface {
	Create(ctx context.Context, name string, dType string, level uint8, index uint8, father *uint64) error
	Delete(ctx context.Context, directory_id uint64) error
	Update(ctx context.Context, directory_id uint64, name string) *DirectoryDTO
	Retrieve(ctx context.Context, level uint8, directory_id *uint64) []*DirectoryDTO
	Move(ctx context.Context, directory_id uint64, father uint64) error
}

type DirectoryVO struct {
	DirectoryId    uint64 `json:"directory_id"`
	DirectoryName  string `json:"directory_name"`
	DirectoryType  string `json:"directory_type"`
	DirectoryLevel uint8  `json:"directory_level"`
	DirectoryIndex uint8  `json:"directory_index"`
}

type DirectoryDTO struct {
	DirectoryId    uint64 `db:"directory_id"`
	DirectoryName  string `db:"directory_name"`
	DirectoryType  string `db:"directory_type"`
	DirectoryLevel uint8  `db:"directory_level"`
	DirectoryIndex uint8  `db:"directory_index"`
}

type DirectoryRelationPO struct {
	Ancestor   uint64 `db:"ancestor"`
	Descendant uint64 `db:"descendant"`
	Distance   uint8  `db:"distance"`
}

// Directory Domain Model
type Directory struct {
	Id             uint64
	DirectoryId    uint64
	DirectoryName  string
	DirectoryType  string
	DirectoryLevel uint8
	DirectoryIndex uint8
	UpdateAt       string
	CreateAt       string
	DeleteAt       string
}

// DirectoryRelation Domain Model
type DirectoryRelation struct {
	Id         uint64
	Ancestor   uint64
	Descendant uint64
	Distance   uint8
	UpdateAt   string
	CreateAt   string
	DeleteAt   string
}
