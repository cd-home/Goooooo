package domain

import "context"

type UserLogicFace interface {
	Register(ctx context.Context, account string, password string) error
}

type UserRepositoryFace interface {
	CreateUserByUserName(ctx context.Context, account string, password string) error
	CreateUserByEmail(ctx context.Context, account string, password string) error
}

// Domain Model And Interface
type User struct {
	Id        int64  `json:"id"`
	UserName  string `json:"user"`
	NickName  string `json:"nickname"`
	Password  string `json:"password"`
	Age       uint8  `json:"age"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	State     uint8  `json:"state"`
	Ip        uint32 `json:"ip"`
	LastLogin string `json:"last_login"`
	UpdateAt  string `json:"update_at"`
	CreateAt  string `json:"create_at"`
	DeleteAt  string `json:"delete_at"`
}
