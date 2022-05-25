package domain

// Role Domain Model
type Role struct {
	Id       uint64
	RoleId   uint64
	RoleName string
}

// RoleRelation Domain Model
type RoleRelation struct {
	Id         uint64
	Ancestor   uint64
	Descendant uint64
	Distance   uint8
	UpdateAt   string
	CreateAt   string
	DeleteAt   string
}
