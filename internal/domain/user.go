package domain

import (
	"context"
	"database/sql"
	"net/http"
)

type UserLogicFace interface {
	Register(ctx context.Context, account string, password string) error
	Login(ctx context.Context, account string, password string) (*UserVO, *UserSession, error)
	SetSession(r *http.Request, rw http.ResponseWriter, sessionObj *UserSession) error
}

type UserRepositoryFace interface {
	CreateUserByUserName(ctx context.Context, account string, password string) error
	CreateUserByEmail(ctx context.Context, account string, password string) error
	CheckAccountExist(ctx context.Context, account string, password string) (*UserDTO, error)
	GetAllUsers(ctx context.Context) ([]*UserDTO, error)
}

type UserEsRepositoryFace interface {
	CreateUserDocument(ctx context.Context, documents *UserEsPO) error
	CreateUserDocuments(ctx context.Context, documents []*UserEsPO) error
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
	State     sql.NullInt16  `db:"state"`
	Ip        sql.NullInt64  `db:"ip"`
	LastLogin sql.NullString `db:"last_login"`
	UpdateAt  string         `db:"update_at"`
	CreateAt  string         `db:"create_at"`
	DeleteAt  sql.NullTime   `db:"delete_at"`
}

// UserEsPO
type UserEsPO struct {
	Id        uint64 `json:"-"`
	UserName  string `json:"username"`
	NickName  string `json:"nickname"`
	Age       uint8  `json:"age"`
	Gender    uint8  `json:"gender"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	State     uint8  `json:"state"`
	Ip        uint64 `json:"ip"`
	LastLogin string `json:"last_login"`
	UpdateAt  string `json:"update_at"`
	CreateAt  string `json:"create_at"`
	DeleteAt  string `json:"delete_at"`
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

// Domain Model
type User struct {
	Id        uint64
	UserName  string
	NickName  string
	Password  string
	Age       uint8
	Gender    string
	Email     string
	Avatar    string
	Phone     string
	State     uint8
	Ip        uint32
	LastLogin string
	UpdateAt  string
	CreateAt  string
	DeleteAt  string
}
