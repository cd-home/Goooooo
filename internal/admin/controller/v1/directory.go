package v1

import (
	"net/http"

	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/pkg/errno"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DirectoryController struct {
	logic domain.DirectoryLogicFace
	log   *zap.Logger
}

func NewDirectoryController(engine *gin.Engine, log *zap.Logger, logic domain.DirectoryLogicFace) {
	ctl := &DirectoryController{
		logic: logic,
		log:   log.WithOptions(zap.Fields(zap.String("module", "DirectoryController"))),
	}
	directory := engine.Group("/api/v1/directory")
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
	resp := types.CommonResponse{Code: 0}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		resp.Message = errno.ParamsParseError
		ctx.JSON(http.StatusOK, resp)
		return
	}
	if err := d.logic.CreateDirectory(
		params.DirectoryName,
		params.DirectoryType,
		params.DirectoryLevel,
		params.DirectoryIndex, params.Father); err != nil {
		resp.Message = errno.Failure
	} else {
		resp.Code = 1
		resp.Message = errno.Success
	}
	ctx.JSON(http.StatusOK, resp)
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
	resp := types.CommonResponse{Code: 0}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.Message = errno.ParamsParseError
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = d.logic.ListDirectory(params.DirectoryLevel, params.Father)
	resp.Code = 1
	resp.Message = errno.Success
	ctx.JSON(http.StatusOK, resp)
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
	resp := types.CommonResponse{Code: 0}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.Message = errno.ParamsParseError
		ctx.JSON(http.StatusOK, resp)
		return
	}
	data := d.logic.RenameDirectory(params.DirectoryId, params.DirectoryName)
	if data == nil {
		resp.Message = errno.Failure
	} else {
		resp.Code = 1
		resp.Message = errno.Success
		resp.Data = data
	}
	ctx.JSON(http.StatusOK, resp)
}

// DeleteDirectory
// @Summary Delete Directory
// @Description Delete Directory
// @Tags Directory
// @Accept  json
// @Produce json
// @Router /delete [PUT]
func (d DirectoryController) DeleteDirectory(ctx *gin.Context) {
	params := types.ListDirectoryParam{}
	resp := types.CommonResponse{Code: 0}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.Message = errno.ParamsParseError
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = d.logic.ListDirectory(params.DirectoryLevel, params.Father)
	resp.Code = 1
	resp.Message = errno.Success
	ctx.JSON(http.StatusOK, resp)
}

// MoveDirectory
// @Summary Move Directory
// @Description Move Directory
// @Tags Directory
// @Accept  json
// @Produce json
// @Router /move [PUT]
func (d DirectoryController) MoveDirectory(ctx *gin.Context) {
	params := types.ListDirectoryParam{}
	resp := types.CommonResponse{Code: 0}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.Message = errno.ParamsParseError
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = d.logic.ListDirectory(params.DirectoryLevel, params.Father)
	resp.Code = 1
	resp.Message = errno.Success
	ctx.JSON(http.StatusOK, resp)
}
