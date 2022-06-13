package types

// CreateRoleParam [Request, POST, Body]
type CreateRoleParam struct {
	RoleName  string  `json:"role_name" binding:"required"`
	RoleLevel uint8   `json:"role_level" binding:"required"`
	RoleIndex uint8   `json:"role_index" binding:"required"`
	Father    *uint64 `json:"father"`
}

// ListDirectoryParam [Request, GET, Query]
type ListRoleParam struct {
	RoleLevel uint8   `form:"role_level" binding:"required"`
	Father    *uint64 `form:"father"`
}

// RenameRoleParam [Request, DELETE, Query]
type DeleteRoleParam struct {
	RoleId uint64 `form:"role_id" binding:"required"`
}

// RenameRoleParam [Request, PUT, Body]
type RenameRoleParam struct {
	RoleId   uint64 `json:"role_id" binding:"required"`
	RoleName string `json:"role_name" binding:"required"`
}

// ListDirectoryParam [Request, POST, Query]
type MoveRoleParam struct {
	RoleId uint64 `json:"role_id" binding:"required"`
	Father uint64 `json:"father" binding:"required"`
}
