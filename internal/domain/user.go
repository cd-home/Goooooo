package domain

import (
	"github.com/GodYao1995/Goooooo/internal/admin/types"
)

type UserLogicFace interface {
	Register(*types.RegisterParam) error
	Login(*types.LoginParam) error
}

type UserRepositoryFace interface {
}

// Domain Model And Interface
type User struct {
	UserName  string `json:"user"`
	NickName  string `json:"nickname"`
	Password  string `json:"password"`
	Age       uint8  `json:"age"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	State     string `json:"state"`
	Ip        uint32 `json:"ip"`
	LastLogin string `json:"last_login"`
	UpdatedAt string `json:"update_at"`
	CreatedAt string `json:"created_at"`
	DeleteAt  string `json:"delete_at"`
}
