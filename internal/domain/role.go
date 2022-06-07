package domain

import "context"

type RoleLogicFace interface {
	CreateRole(ctx context.Context)
}

// RoleRepositoryFace
type RoleRepositoryFace interface {
	CreateRole(ctx context.Context) error
	CreateRelation(ctx context.Context) error
	Delete(ctx context.Context)
	Update(ctx context.Context)
	Retrieve(ctx context.Context)
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

// RolePO [Write To DB, For Insert]
type RolePO struct {
	RoleId    uint64 `db:"role_id"`
	RoleName  string `db:"role_name"`
	RoleLevel uint8  `db:"role_level"`
	RoleIndex uint8  `db:"role_index"`
}
