package v1

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/internal/pkg/session"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type UserLogic struct {
	repo  domain.UserRepositoryFace
	store *session.RedisStore
}

func NewUserLogic(repo domain.UserRepositoryFace, store *session.RedisStore) domain.UserLogicFace {
	return &UserLogic{repo: repo, store: store}
}

// Register
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

// Login
func (logic *UserLogic) Login(ctx context.Context, account string, password string) (*domain.UserVO, *domain.UserSession, error) {
	user, err := logic.repo.CheckAccountExist(ctx, account, password)
	if user == nil {
		return nil, nil, err
	}
	// Check Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, nil, errno.ErrorUserPassword
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

// SetSession
func (logic *UserLogic) SetSession(r *http.Request, rw http.ResponseWriter, obj *domain.UserSession) error {
	session, _ := logic.store.Get(r, "SESSIONID")
	// store session
	values, _ := json.Marshal(obj)
	session.Values["user"] = values
	// TODO 后期修改到配置项
	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 60 * 60 * 2,
		// 需要false 否则前端无法读取Cookie
		HttpOnly: false,
		Secure:   false,
	}
	// write Cookie and store to Redis Session
	if err := session.Save(r, rw); err != nil {
		return err
	}
	return nil
}
