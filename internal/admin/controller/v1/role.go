package v1

import (
	"net/http"

	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/admin/version"
	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/internal/pkg/session"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RoleController struct {
	logic domain.RoleLogicFace
	log   *zap.Logger
	store *session.RedisStore
	perm  *casbin.Enforcer
}

func NewRoleController(apiV1 *version.APIV1, log *zap.Logger, logic domain.RoleLogicFace, store *session.RedisStore, perm *casbin.Enforcer) {
	ctl := &RoleController{
		logic: logic,
		log:   log.WithOptions(zap.Fields(zap.String("module", "RoleController"))),
		store: store,
		perm:  perm,
	}
	// API version
	// v1 := apiV1.Group.Group("/role").Use(auth.AuthMiddleware(store)).Use(permission.PermissionMiddleware(perm))
	v1 := apiV1.Group.Group("/role")
	{
		v1.POST("/create", ctl.CreateRole)
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
	err = r.logic.CreateRole(ctx, params.RoleName, params.RoleLevel, params.RoleIndex, params.Parent)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Message = errno.RoleCreateSuccess
		resp.Code = 0
	}
	ctx.JSON(http.StatusOK, resp)
}
