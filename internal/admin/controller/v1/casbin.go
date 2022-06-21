package v1

import (
	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/admin/version"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/internal/pkg/middleware/auth"
	"github.com/GodYao1995/Goooooo/internal/pkg/middleware/permission"
	"github.com/GodYao1995/Goooooo/internal/pkg/res"
	"github.com/GodYao1995/Goooooo/internal/pkg/session"
	"github.com/GodYao1995/Goooooo/pkg/xhttp"
	"github.com/GodYao1995/Goooooo/pkg/xtracer"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type CasbinController struct {
	perm *casbin.Enforcer
}

func NewCasbinController(apiV1 *version.APIV1, store *session.RedisStore, perm *casbin.Enforcer, xtracer *xtracer.XTracer) {
	v1 := apiV1.Group.Group("/perm").Use(auth.AuthMiddleware(store)).Use(permission.PermissionMiddleware(perm))
	ctl := &CasbinController{perm: perm}
	{
		v1.POST("/create", ctl.CreatePermission)
		v1.POST("/creates", ctl.CreatePermissions)
	}
}

// CreatePermissions
// @Summary Create Permissions
// @Description Create Permissions
// @Tags Permission
// @Accept  json
// @Produce json
// @Param
// @Success 0 {object} types.CommonResponse {"code":1,"data":null,"msg":"Success"}
// @Failure 1 {object} types.CommonResponse {"code":1,"data":null,"msg":"Error"}
// @Router /create [POST]
func (c CasbinController) CreatePermission(ctx *gin.Context) {
	resp := res.CommonResponse{Code: 1}
	params := types.CreatePermissionParam{}
	if ok, valid := xhttp.ShouldBindJSON(ctx, &params); !ok {
		resp.Message = valid
		resp.Failure(ctx)
		return
	}
	if params.Ptype == "p" {
		// 用户权限
		if ok, err := c.perm.AddPolicy(params.Subject, params.Object, params.Action, params.Version); !ok {
			resp.Message = err.Error()
			resp.Failure(ctx)
			return
		}
	} else if params.Ptype == "g" {
		// 角色权限
		if ok, err := c.perm.AddGroupingPolicy(params.Subject, params.Object, params.Action, params.Version); !ok {
			resp.Message = err.Error()
			resp.Failure(ctx)
			return
		}
	} else {
		resp.Message = errno.ErrorParamsParse
		resp.Failure(ctx)
		return
	}
	resp.Code = 0
	resp.Message = errno.CreatePermissionSuccess
	resp.Success(ctx)
}

// CreatePermission
// @Summary Create Permission
// @Description Create Permission
// @Tags Permission
// @Accept  json
// @Produce json
// @Param
// @Success 0 {object} types.CommonResponse {"code":1,"data":null,"msg":"Success"}
// @Failure 1 {object} types.CommonResponse {"code":1,"data":null,"msg":"Error"}
// @Router /create [POST]
func (c CasbinController) CreatePermissions(ctx *gin.Context) {
	// 用户权限
	c.perm.AddPolicy()
	// 角色权限
	c.perm.AddGroupingPolicy()
}
