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
		log:   log.WithOptions(zap.Fields(zap.String("module", "UserController"))),
	}
	directory := engine.Group("/api/v1/directory")
	{
		directory.POST("/create", ctl.CreateDirectory)
		directory.GET("/list", ctl.ListDirectory)
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
	common := types.CommonResponse{Code: 0}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		common.Message = errno.ParamsParseError
		ctx.JSON(http.StatusOK, common)
		return
	}
	err := d.logic.CreateDirectory(
		params.DirectoryName,
		params.DirectoryType,
		params.DirectoryLevel,
		params.DirectoryIndex, params.Father)
	if err != nil {
		common.Message = errno.Failure
		ctx.JSON(200, common)
		return
	} else {
		common.Code = 0
		common.Message = errno.Success
		ctx.JSON(200, common)
		return
	}
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
	common := types.CommonResponse{Code: 0}
	if err := ctx.ShouldBind(&params); err != nil {
		common.Message = err.Error()
		ctx.JSON(http.StatusOK, common)
		return
	}
	data := d.logic.ListDirectory(params.DirectoryLevel, params.Father)
	common.Code = 0
	common.Message = errno.Success
	common.Data = data
	ctx.JSON(200, common)
}
