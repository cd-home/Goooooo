package v1

import (
	"net/http"

	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/admin/version"
	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/internal/pkg/middleware/auth"
	"github.com/GodYao1995/Goooooo/internal/pkg/session"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RoleController struct {
	logic domain.RoleLogicFace
	log   *zap.Logger
}

func NewRoleController(apiV1 *version.APIV1, log *zap.Logger, logic domain.RoleLogicFace, store *session.RedisStore, perm *casbin.Enforcer) {
	ctl := &RoleController{
		logic: logic,
		log:   log.WithOptions(zap.Fields(zap.String("module", "RoleController"))),
	}
	// API version
	v1 := apiV1.Group.Group("/role")
	needAuth := v1.Use(auth.AuthMiddleware(store))
	{
		needAuth.GET("/list", ctl.ListRole)
	}
	// needPerm := needAuth.Use(permission.PermissionMiddleware(perm))
	{
		needAuth.POST("/create", ctl.CreateRole)
		needAuth.POST("/delete", ctl.CreateRole)
	}
}

// CreateRole
// @Summary Create Role
// @Description Create Role
// @Tags Role
// @Accept  json
// @Produce json
// @Router /create [POST]
func (r RoleController) CreateRole(ctx *gin.Context) {
	params := types.CreateRoleParam{}
	resp := types.CommonResponse{Code: 1}
	var err error
	if err = ctx.ShouldBindJSON(&params); err != nil {
		resp.Message = errno.ErrorParamsParse.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	err = r.logic.CreateRole(ctx, params.RoleName, params.RoleLevel, params.RoleIndex, params.Father)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Message = errno.RoleCreateSuccess
		resp.Code = 0
	}
	ctx.JSON(http.StatusOK, resp)
}

// DeleteRole
// @Summary Delete Role
// @Description Delete Role
// @Tags Role
// @Accept  json
// @Produce json
// @Router /delete [POST]
func (r RoleController) DeleteRole(ctx *gin.Context) {
	resp := types.CommonResponse{Code: 1}
	ctx.JSON(http.StatusOK, resp)
}

// ListRole
// @Summary List Role
// @Description List Role
// @Tags Role
// @Accept  json
// @Produce json
// @Router /list [GET]
func (r RoleController) ListRole(ctx *gin.Context) {
	params := types.ListRoleParam{}
	resp := types.CommonResponse{Code: 1}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.Message = errno.ErrorParamsParse.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	views, err := r.logic.Retrieve(ctx, params.RoleLevel, params.Father)
	if err != nil {
		resp.Message = err.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp.Message = errno.Success
	resp.Data = views
	ctx.JSON(http.StatusOK, resp)
}
