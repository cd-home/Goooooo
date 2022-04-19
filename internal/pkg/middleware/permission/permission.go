package permission

import (
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// just for testing purposes
		subject := ctx.Query("sub")
		object := ctx.Query("obj")
		action := ctx.Query("act")
		version := strings.Split(ctx.Request.URL.Path, "/")[2]
		ok, _ := e.Enforce(subject, object, action, version)
		if !ok {
			ctx.JSON(401, map[string]interface{}{
				"message": "No permission",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
