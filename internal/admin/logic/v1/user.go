package v1

import (
	"context"
	"errors"

	"github.com/GodYao1995/Goooooo/internal/domain"
)

type UserLogic struct {
	repo domain.UserRepositoryFace
}

func NewUserLogic(repo domain.UserRepositoryFace) domain.UserLogicFace {
	return &UserLogic{repo: repo}
}

func (logic *UserLogic) Register(ctx context.Context, account string, password string) error {
	err := logic.repo.CreateUserByUserName(ctx, account, password)
	if err != nil {
		return errors.New("注册失败, " + err.Error())
	}
	return nil
}

func (logic *UserLogic) Login() error {
	return nil
}
