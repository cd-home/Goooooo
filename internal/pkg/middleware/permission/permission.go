package permission

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Login User
		v, _ := ctx.Get("user")
		session := v.(domain.UserSession)

		parts := strings.Split(ctx.Request.URL.Path, "/")

		userSub := session.Id
		roleSub := session.Role

		object := "/" + parts[3] + "/" + parts[4]
		action := parts[5]
		version := parts[2]

		log.Println(userSub, object, action, version)
		log.Println(roleSub, object, action, version)

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
			resp := types.CommonResponse{Code: 1}
			resp.Message = errno.NoPermission
			ctx.JSON(http.StatusOK, resp)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
