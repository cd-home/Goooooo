package v1

import (
	"net/http"
	"strconv"

	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/admin/version"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/pkg/tools"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileController struct {
	log    *zap.Logger
	file   string
	upload string
}

func NewFileController(apiV1 *version.APIV1, log *zap.Logger) {
	ctl := &FileController{
		log:    log.WithOptions(zap.Fields(zap.String("module", "FileController"))),
		file:   "file",
		upload: "../upload/",
	}
	v1 := apiV1.Group
	file := v1.Group("/file")
	{
		file.GET("/list", ctl.ListFile)
		file.POST("/upload", ctl.UploadLocal)
		file.POST("/oss", ctl.UploadOss)
	}
}

// ListFile
// @Summary List File
// @Description List File
// @Tags File
// @Accept  json
// @Produce json
// @Router /list [GET]
func (d FileController) ListFile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

// UploadLocal
// @Summary UploadLocal
// @Description List File
// @Tags File
// @Accept  json
// @Produce json
// @Router /upload [POST]
func (d FileController) UploadLocal(ctx *gin.Context) {
	resp := types.CommonResponse{Code: 1}
	f, err := ctx.FormFile(d.file)
	if err != nil {
		resp.Message = errno.ErrorUploadFile.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	target := d.upload + f.Filename + strconv.Itoa(int(tools.SnowId()))
	if err = ctx.SaveUploadedFile(f, target); err != nil {
		resp.Message = errno.ErrorUploadFile.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp.Code = 0
	resp.Message = errno.UploadSuccess
	ctx.JSON(http.StatusOK, resp)
}

// UploadOss
// @Summary UploadOss
// @Description Upload Oss
// @Tags File
// @Accept  json
// @Produce json
// @Router /oss [POST]
func (d FileController) UploadOss(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
