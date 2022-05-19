package domain

import (
	"context"
	"database/sql"
)

type UserLogicFace interface {
	Register(ctx context.Context, account string, password string) error
	Login(ctx context.Context, account string, password string) (*UserVO, *UserSession, error)
}

type UserRepositoryFace interface {
	CreateUserByUserName(ctx context.Context, account string, password string) error
	CreateUserByEmail(ctx context.Context, account string, password string) error
	CheckAccountExist(ctx context.Context, account string, password string) (*UserDTO, error)
}

// User DTO
type UserDTO struct {
	Id        uint64         `db:"id"`
	UserName  string         `db:"username"`
	NickName  sql.NullString `db:"nickname"`
	Password  string         `db:"password"`
	Age       sql.NullInt16  `db:"age"`
	Gender    sql.NullInt16  `db:"gender"`
	Email     sql.NullString `db:"email"`
	Avatar    sql.NullString `db:"avatar"`
	Phone     sql.NullString `db:"phone"`
	State     sql.NullString `db:"state"`
	Ip        sql.NullInt64  `db:"ip"`
	LastLogin sql.NullString `db:"last_login"`
	UpdateAt  string         `db:"update_at"`
	CreateAt  string         `db:"create_at"`
	DeleteAt  sql.NullTime   `db:"delete_at"`
}

// User VO
type UserVO struct {
	Id        uint64 `json:"-"`
	UserName  string `json:"username"`
	NickName  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	LastLogin string `json:"last_login"`
}

// User Session
type UserSession struct {
	Id        uint64 `json:"id"`
	UserName  string `json:"username"`
	NickName  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	LastLogin string `json:"last_login"`
}

// Domain Model And Interface
type User struct {
	Id        uint64 `json:"id"`
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
