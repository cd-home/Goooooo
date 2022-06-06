package permission

import (
	"net/http"
	"strings"

	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// just for testing purposes
		subject := ctx.Query("sub")
		object := ctx.Query("obj")
		action := ctx.Query("act")
		resp := types.CommonResponse{Code: 1}
		version := strings.Split(ctx.Request.URL.Path, "/")[2]
		ok, _ := e.Enforce(subject, object, action, version)
		if !ok {
			resp.Message = errno.NoPermission
			ctx.JSON(http.StatusOK, resp)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
