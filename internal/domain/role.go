package domain

import "context"

type RoleLogicFace interface {
	CreateRole(ctx context.Context, roleName string, roleLevel uint8, roleIndex uint8, parent *uint64) error
	DeleteRole(ctx context.Context, roleId uint64) error
	UpdateRole(ctx context.Context, roleId uint64, roleName string) error
	RetrieveRoles(ctx context.Context, roleLevel uint8, father *uint64) ([]*RoleEntityVO, error)
}

// RoleRepositoryFace
type RoleRepositoryFace interface {
	Create(ctx context.Context, roleId uint64, roleName string, roleLevel uint8, roleIndex uint8, fa *uint64) error
	Delete(ctx context.Context, roleId uint64) error
	Update(ctx context.Context, roleId uint64, roleName string) error
	Retrieve(ctx context.Context, roleLevel uint8, father *uint64) ([]*RoleEntityDTO, error)
}

// RoleEntity [Mapping To DB Fields, For Select]
type RoleEntity struct {
	Id        uint64 `db:"id"`
	RoleId    uint64 `db:"role_id"`
	RoleName  string `db:"role_name"`
	RoleLevel uint8  `db:"role_level"`
	RoleIndex uint8  `db:"role_index"`
	UpdateAt  string `db:"update_at"`
	CreateAt  string `db:"create_at"`
	DeleteAt  string `db:"delete_at"`
}

// RoleRelationEntity [Mapping To DB Fields, For Select]
type RoleRelationEntity struct {
	Id         uint64 `db:"id"`
	Ancestor   uint64 `db:"ancestor"`
	Descendant uint64 `db:"descendant"`
	Distance   uint8  `db:"distance"`
	UpdateAt   string `db:"update_at"`
	CreateAt   string `db:"create_at"`
	DeleteAt   string `db:"delete_at"`
}

// RoleRelationPO [For Write DB]
type RoleRelationPO struct {
	Ancestor   uint64 `db:"ancestor"`
	Descendant uint64 `db:"descendant"`
	Distance   uint8  `db:"distance"`
}

// RoleEntityDTO
type RoleEntityDTO struct {
	RoleId    uint64 `db:"role_id"`
	RoleName  string `db:"role_name"`
	RoleLevel uint8  `db:"role_level"`
	RoleIndex uint8  `db:"role_index"`
	UpdateAt  string `db:"update_at"`
	CreateAt  string `db:"create_at"`
}

// RoleEntityVO
type RoleEntityVO struct {
	RoleId    uint64 `json:"role_id"`
	RoleName  string `json:"role_name"`
	RoleLevel uint8  `json:"role_level"`
	RoleIndex uint8  `json:"role_index"`
	UpdateAt  string `json:"update_at"`
	CreateAt  string `json:"create_at"`
}
