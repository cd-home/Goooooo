package v1

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
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

type UserRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewUserRepository(db *sqlx.DB, log *zap.Logger) domain.UserRepositoryFace {
	return &UserRepository{db: db, log: log.WithOptions(zap.Fields(zap.String("module", "UserRepository")))}
}

func (repo *UserRepository) CreateUserByUserName(ctx context.Context, account string, password string) error {
	// check if user already exist
	var user User
	var err error
	err = repo.db.Get(&user, `SELECT id FROM user WHERE username = ?`, account)
	if err != nil {
		repo.log.WithOptions(zap.Fields(zap.String("Repo Get", "CreateUserByUserName"))).Info(err.Error())
		return errors.New("操作失败, 请重试")
	}
	repo.log.WithOptions(zap.Fields(zap.String("Repo", "CreateUserByUserName"))).Debug(fmt.Sprint(user.UserName))
	if err == nil {
		return errors.New("该用户名已经被注册")
	}
	// create user by username
	bcryptPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err = repo.db.Exec(`INSERT INTO user (username, password) VALUES (?, ?)`, account, bcryptPwd)
	if err != nil {
		repo.log.WithOptions(zap.Fields(zap.String("Repo Exec", "CreateUserByUserName"))).Info(err.Error())
		return errors.New("操作失败, 请重试")
	}
	return nil
}
func (repo *UserRepository) CreateUserByEmail(ctx context.Context, account string, password string) error {
	bcryptPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err := repo.db.Exec(`INSERT INTO user (email, password) VALUES (?, ?)`, account, bcryptPwd)
	return err
}
