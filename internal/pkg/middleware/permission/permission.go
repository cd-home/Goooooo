package permission

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/consts"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/internal/pkg/res"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Login User
		v, _ := ctx.Get(consts.SROREKEY)
		session := v.(domain.UserSession)

		parts := strings.Split(ctx.Request.URL.Path, "/")

		userSub := session.Id
		roleSub := session.Role

		// TODO: Optimization Permission
		object := parts[3]
		action := parts[4]
		version := parts[2]

		// Remind: Casbin sub Type Must String
		ok, _ := e.Enforce(strconv.Itoa(int(userSub)), object, action, version)

		// User In Role
		var okk = true
		if len(roleSub) > 0 {
			for _, v := range roleSub {
				_ok, _ := e.Enforce(strconv.Itoa(int(v)), object, action, version)
				if !_ok {
					okk = false
					break
				}
			}
		}
		if !ok && !okk {
			resp := res.CommonResponse{Code: 1}
			resp.Message = errno.NoPermission
			ctx.JSON(http.StatusOK, resp)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
