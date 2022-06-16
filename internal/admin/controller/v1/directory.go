package v1

import (
	"context"
	"net/http"

	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/admin/version"
	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/internal/pkg/middleware/auth"
	"github.com/GodYao1995/Goooooo/internal/pkg/middleware/permission"
	"github.com/GodYao1995/Goooooo/internal/pkg/res"
	"github.com/GodYao1995/Goooooo/internal/pkg/session"
	"github.com/GodYao1995/Goooooo/pkg/xhttp/param"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type DirectoryController struct {
	logic domain.DirectoryLogicFace
	log   *zap.Logger
}

func NewDirectoryController(apiV1 *version.APIV1, log *zap.Logger,
	logic domain.DirectoryLogicFace, store *session.RedisStore, perm *casbin.Enforcer) {
	ctl := &DirectoryController{
		logic: logic,
		log:   log.WithOptions(zap.Fields(zap.String("module", "DirectoryController"))),
	}
	v1 := apiV1.Group
	directory := v1.Group("/directory").Use(auth.AuthMiddleware(store)).Use(permission.PermissionMiddleware(perm))
	{
		directory.POST("/create", ctl.CreateDirectory)
		directory.GET("/list", ctl.ListDirectory)
		directory.PUT("/move", ctl.MoveDirectory)
		directory.PUT("/rename", ctl.RenameDirectory)
		directory.DELETE("/delete", ctl.DeleteDirectory)
	}
}

// CreateDirectory
// @Summary Create Directory
// @Description Create Directory
// @Tags Directory
// @Accept  json
// @Produce json
// @Router /create [POST]
func (d DirectoryController) CreateDirectory(ctx *gin.Context) {
	params := types.CreateDirectoryParam{}
	resp := res.CommonResponse{Code: 1}
	if ok, valid := param.ShouldBindJSON(ctx, &params); !ok {
		resp.Message = valid
		resp.Failure(ctx)
		return
	}
	if err := d.logic.CreateDirectory(
		ctx,
		params.DirectoryName,
		params.DirectoryType,
		params.DirectoryLevel,
		params.DirectoryIndex, params.Father); err != nil {
		resp.Message = errno.Failure
	} else {
		resp.Code = 0
		resp.Message = errno.Success
	}
	resp.Success(ctx)
}

// ListDirectory
// @Summary List Directory
// @Description List Directory
// @Tags Directory
// @Accept  json
// @Produce json
// @Router /list [POST]
func (d DirectoryController) ListDirectory(ctx *gin.Context) {
	params := types.ListDirectoryParam{}
	resp := res.CommonResponse{Code: 1}
	if ok, valid := param.ShouldBind(ctx, &params); !ok {
		resp.Message = valid
		resp.Failure(ctx)
		return
	}
	resp.Data = d.logic.ListDirectory(ctx, params.DirectoryLevel, params.Father)
	resp.Code = 0
	resp.Message = errno.Success
	resp.Success(ctx)
}

// RenameDirectory
// @Summary Rename Directory
// @Description Rename Directory
// @Tags Directory
// @Accept  json
// @Produce json
// @Router /rename [PUT]
func (d DirectoryController) RenameDirectory(ctx *gin.Context) {
	params := types.RenameDirectoryParam{}
	resp := res.CommonResponse{Code: 1}
	if ok, valid := param.ShouldBindJSON(ctx, &params); !ok {
		resp.Message = valid
		resp.Failure(ctx)
	}
	data := d.logic.RenameDirectory(ctx, params.DirectoryId, params.DirectoryName)
	if data == nil {
		resp.Message = errno.Failure
	} else {
		resp.Code = 0
		resp.Message = errno.Success
		resp.Data = data
	}
	resp.Success(ctx)
}

// DeleteDirectory
// @Summary Delete Directory
// @Description Delete Directory
// @Tags Directory
// @Accept  json
// @Produce json
// @Router /delete [DELETE]
func (d DirectoryController) DeleteDirectory(ctx *gin.Context) {
	// TODO: delete directory
	params := types.ListDirectoryParam{}
	resp := res.CommonResponse{Code: 1}
	if ok, valid := param.ShouldBindQuery(ctx, &params); !ok {
		resp.Message = valid
		resp.Failure(ctx)
		return
	}
	resp.Data = d.logic.ListDirectory(ctx, params.DirectoryLevel, params.Father)
	resp.Code = 0
	resp.Message = errno.Success
	ctx.JSON(http.StatusOK, resp)
}

// MoveDirectory
// @Summary Move Directory
// @Description Move Directory
// @Tags Directory
// @Accept  json
// @Produce json
// @Router /move [POST]
func (d DirectoryController) MoveDirectory(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "DirectoryController-MoveDirectory")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("DirectoryController", "MoveDirectory")
		span.Finish()
	}()
	params := types.MoveDirectoryParam{}
	resp := res.CommonResponse{Code: 1}
	if ok, valid := param.ShouldBind(ctx, &params); !ok {
		resp.Message = valid
		resp.Failure(ctx)
		return
	}
	err := d.logic.MoveDirectory(next, params.DirectoryId, params.Father)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = errno.Success
	}
	resp.Success(ctx)
}
