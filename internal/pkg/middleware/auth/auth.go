package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cd-home/Goooooo/internal/domain"
	"github.com/cd-home/Goooooo/internal/pkg/consts"
	"github.com/cd-home/Goooooo/internal/pkg/errno"
	"github.com/cd-home/Goooooo/internal/pkg/res"
	"github.com/cd-home/Goooooo/internal/pkg/session"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(store *session.RedisStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := res.CommonResponse{Code: 1}
		var user domain.UserSession
		session, err := store.New(ctx.Request, consts.SESSIONID)
		if errors.Is(err, errno.ErrorRedisEmpty) || errors.Is(err, http.ErrNoCookie) {
			resp.Message = errno.ErrorUserNotLogin.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}
		err = json.Unmarshal(session.Values[consts.SROREKEY].([]byte), &user)
		if err != nil {
			ctx.Abort()
			return
		}
		ctx.Set(consts.SROREKEY, user)
		ctx.Next()
	}
}
