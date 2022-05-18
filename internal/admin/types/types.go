package types

// CommonResponse [Response]
type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// RegisterParam [Request, POST, Body]
type RegisterParam struct {
	Account  string `json:"account" binding:"required,min=4,max=50"`
	Password string `json:"password" binding:"required,min=6,max=18"`
}

type LoginParam struct {
	Method   string `json:"method"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

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
