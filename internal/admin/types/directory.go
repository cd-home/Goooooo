package types

// CreateDirectoryParam [Request, POST, Body]
type CreateDirectoryParam struct {
	DirectoryName  string  `json:"directory_name" binding:"required"`
	DirectoryType  string  `json:"directory_type" binding:"required"`
	DirectoryLevel uint8   `json:"directory_level" binding:"required"`
	DirectoryIndex uint8   `json:"directory_index" binding:"required"`
	Father         *uint64 `json:"father"`
}

// ListDirectoryParam [Request, GET, Query]
type ListDirectoryParam struct {
	DirectoryLevel uint8   `form:"directory_level" binding:"required"`
	Father         *uint64 `form:"father"`
}

// RenameDirectoryParam [Request, POST, Body]
type RenameDirectoryParam struct {
	DirectoryId   uint64 `json:"directory_id" binding:"required"`
	DirectoryName string `json:"directory_name" binding:"required"`
}

