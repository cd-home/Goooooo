package domain

type DirectoryLogicFace interface {
	CreateDirectory(name string, dType string, level uint8, index uint8, father *uint64) error
	ListDirectory(level uint8, directory_id *uint64) []*DirectoryVO
}

type DirectoryRepositoryFace interface {
	CreateDirectory(name string, dType string, level uint8, index uint8, father *uint64) error
	ListDirectory(level uint8, directory_id *uint64) []*DirectoryDTO
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

type DirectoryRelation struct {
	Ancestor   uint64 `db:"ancestor"`
	Descendant uint64 `db:"descendant"`
	Distance   uint8  `db:"distance"`
}
