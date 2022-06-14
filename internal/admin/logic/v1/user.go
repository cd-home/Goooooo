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
	"github.com/opentracing/opentracing-go"
	"golang.org/x/crypto/bcrypt"
)

type UserLogic struct {
	repo  domain.UserRepositoryFace
	store *session.RedisStore
	espo  domain.UserEsRepositoryFace
}

func NewUserLogic(repo domain.UserRepositoryFace, store *session.RedisStore, espo domain.UserEsRepositoryFace) domain.UserLogicFace {
	return &UserLogic{repo: repo, store: store, espo: espo}
}

// Register
func (logic *UserLogic) Register(ctx context.Context, account string, password string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "UserLogic-Register")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("UserLogic", "Register")
		span.Finish()
	}()
	user, err := logic.repo.RetrieveByUserName(next, account, password)
	// DataBaseError
	if user == nil && errors.Is(err, errno.ErrorDataBase) {
		return err
	}
	// User exist
	if user != nil && errors.Is(err, errno.ErrorUserRecordExist) {
		return err
	}
	err = logic.repo.CreateByUserName(next, account, password)
	if err != nil {
		return err
	}
	return nil
}

// Login
func (logic *UserLogic) Login(ctx context.Context,
	r *http.Request, rw http.ResponseWriter, account string, password string) (*domain.UserVO, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "UserLogic-Login")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("UserLogic", "Login")
		span.Finish()
	}()
	//  CheckAccountExist
	user, err := logic.repo.RetrieveByUserName(next, account, password)
	if user == nil {
		return nil, err
	}

	// Check Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errno.ErrorUserPassword
	}

	Vo := &domain.UserVO{
		Id:       user.Id,
		UserName: user.UserName,
	}

	// GetRoleByUserId
	roles, err := logic.repo.RetrieveRoleByUserId(next, user.Id)
	if err != nil {
		return nil, err
	}

	// store session
	sessionObj := &domain.UserSession{
		Id:       user.Id,
		UserName: user.UserName,
		Role:     roles,
	}

	session, _ := logic.store.Get(r, "SESSIONID")
	values, _ := json.Marshal(sessionObj)
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
		return nil, err
	}
	return Vo, nil
}
