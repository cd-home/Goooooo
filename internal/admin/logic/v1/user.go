package v1

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/consts"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/internal/pkg/session"
	"github.com/GodYao1995/Goooooo/pkg/tools"
	"github.com/GodYao1995/Goooooo/pkg/xtime"
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
	// 不存在才可以注册
	_, err := logic.repo.RetrieveByUserName(next, account, password)
	if !errors.Is(err, errno.ErrorUserRecordNotExist) {
		return err
	}
	err = logic.repo.CreateByUserName(next, account, password)
	if err != nil {
		return err
	}
	return nil
}

// Login
func (logic *UserLogic) Login(ctx context.Context, r *http.Request, rw http.ResponseWriter, account string, password string) (*domain.UserVO, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "UserLogic-Login")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("UserLogic", "Login")
		span.Finish()
	}()
	//  CheckAccountExist
	user, err := logic.repo.RetrieveByUserName(next, account, password)
	if !errors.Is(err, errno.ErrorUserRecordExist) {
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
	session, _ := logic.store.Get(r, consts.SESSIONID)
	values, _ := json.Marshal(sessionObj)
	session.Values[consts.SROREKEY] = values
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

// RetrieveAllUser
func (logic UserLogic) RetrieveAllUser(ctx context.Context) ([]*domain.UserVO, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "UserLogic-RetrieveAllUser")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("UserLogic", "RetrieveAllUser")
		span.Finish()
	}()
	dtos, err := logic.repo.RetrieveAllUsers(next)
	if err != nil {
		return nil, err
	}
	vos := make([]*domain.UserVO, 0)
	for _, obj := range dtos {
		var (
			nickname *string
			age      *uint8
			gender   *uint8
			state    *uint8
			ip       *string
		)
		if obj.NickName.Valid {
			nickname = &obj.NickName.String
		}
		if obj.Age.Valid {
			_age := uint8(obj.Age.Int16)
			age = &_age
		}
		if obj.Gender.Valid {
			_gender := uint8(obj.Gender.Int16)
			gender = &_gender
		}
		if obj.State.Valid {
			_state := uint8(obj.State.Int16)
			state = &_state
		}
		if obj.Ip.Valid {
			_ip := tools.UintIpToString(uint32(obj.Ip.Int64))
			ip = &_ip
		}
		// TODO 后续优化
		vos = append(vos, &domain.UserVO{
			UserName:  obj.UserName,
			NickName:  nickname,
			Age:       age,
			Gender:    gender,
			Email:     obj.Email.String,
			Avatar:    obj.Avatar.String,
			Phone:     obj.Phone.String,
			State:     state,
			Ip:        ip,
			LastLogin: obj.LastLogin.String,
			UpdateAt:  obj.UpdateAt,
			DeleteAt:  obj.DeleteAt.Time.Format(xtime.TimeLayOut),
		})
	}
	return vos, nil
}

// ModifyPassword
func (logic UserLogic) ModifyPassword(ctx context.Context, originPassword, newPassword string, uid uint64) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "UserLogic-ModifyPassword")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("UserLogic", "ModifyPassword")
		span.Finish()
	}()
	// check user
	user, err := logic.repo.RetrieveByUserId(next, uid)
	if err != nil {
		return err
	}
	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(originPassword))
	if err != nil {
		return errno.ErrorOriginPassword
	}
	return logic.repo.ModifyPassword(next, originPassword, newPassword, uid)
}

// Logout
func (logic UserLogic) Logout(ctx context.Context, r *http.Request, w http.ResponseWriter) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "UserLogic-Logout")
	defer func() {
		span.SetTag("UserLogic", "Logout")
		span.Finish()
	}()
	session, err := logic.store.Get(r, consts.SESSIONID)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	if err = session.Save(r, w); err != nil {
		return err
	}
	return nil
}
