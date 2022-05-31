package v1

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewUserRepository(db *sqlx.DB, log *zap.Logger) domain.UserRepositoryFace {
	return &UserRepository{
		db:  db,
		log: log.WithOptions(zap.Fields(zap.String("module", "UserRepository"))),
	}
}

// CreateUserByUserName
func (repo *UserRepository) CreateUserByUserName(ctx context.Context, account string, password string) error {
	local := zap.Fields(zap.String("Repo", "CreateUserByUserName"))
	// create user by username
	bcryptPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err := repo.db.Exec(`INSERT INTO user (username, password) VALUES (?, ?)`, account, string(bcryptPwd))
	if err != nil {
		repo.log.WithOptions(local).Info(err.Error())
		return err
	}
	logger := fmt.Sprint(account, " Register At ", time.Now().Local().Format("2006-01-02 15:04:05"))
	repo.log.WithOptions(local).Debug(logger)
	return nil
}

// CreateUserByEmail
func (repo *UserRepository) CreateUserByEmail(ctx context.Context, account string, password string) error {
	bcryptPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err := repo.db.Exec(`INSERT INTO user (email, password) VALUES (?, ?)`, account, bcryptPwd)
	return err
}

// CheckAccountExist
func (repo *UserRepository) CheckAccountExist(ctx context.Context, account string, password string) (*domain.UserDTO, error) {
	// check if user already exist
	var user domain.UserDTO
	var err error
	local := zap.Fields(zap.String("Repo", "CheckAccountExist"))
	err = repo.db.Get(&user, `
		SELECT 
			id, username, nickname, password, create_at 
		FROM user WHERE username = ? AND delete_at is null`, account)
	// database error
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		repo.log.WithOptions(local).Info(err.Error())
		return nil, errno.ErrorDataBase
	}
	// user already exist
	if err == nil {
		logger := fmt.Sprint(user.UserName, " Registered At ", user.CreateAt)
		repo.log.WithOptions(local).Debug(logger)
		return &user, errno.ErrorUserRecordExist
		// use not existing
	} else {
		return nil, errno.ErrorUserRecordNotExist
	}
}

// GetAllUsers
func (repo *UserRepository) GetAllUsers(ctx context.Context) ([]*domain.UserDTO, error) {
	var err error
	var users []*domain.UserDTO
	err = repo.db.Select(&users, `SELECT * FROM user`)
	return users, err
}
