package v1

import (
	"context"
	"errors"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/internal/pkg/session"
	"golang.org/x/crypto/bcrypt"
)

type UserLogic struct {
	repo    domain.UserRepositoryFace
	session *session.RedisStore
}

func NewUserLogic(repo domain.UserRepositoryFace, session *session.RedisStore) domain.UserLogicFace {
	return &UserLogic{repo: repo, session: session}
}

func (logic *UserLogic) Register(ctx context.Context, account string, password string) error {
	user, err := logic.repo.CheckAccountExist(ctx, account, password)
	// DataBaseError
	if user == nil && errors.Is(err, errno.ErrorDataBase) {
		return err
	}
	// User exist
	if user != nil && errors.Is(err, errno.ErrorUserRecordExist) {
		return err
	}
	err = logic.repo.CreateUserByUserName(ctx, account, password)
	if err != nil {
		return err
	}
	return nil
}

func (logic *UserLogic) Login(ctx context.Context, account string, password string) (*domain.UserVO, *domain.UserSession, error) {
	user, err := logic.repo.CheckAccountExist(ctx, account, password)
	if user == nil {
		return nil, nil, err
	}
	// Check Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, nil, err
	}
	obj := &domain.UserVO{
		Id:       user.Id,
		UserName: user.UserName,
	}
	sessions := &domain.UserSession{
		Id:       user.Id,
		UserName: user.UserName,
	}
	return obj, sessions, nil

}
