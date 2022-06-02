package v1

import (
	"net/http"

	"github.com/GodYao1995/Goooooo/internal/admin/version"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileController struct {
	log *zap.Logger
}

func NewFileController(apiV1 *version.APIV1, log *zap.Logger) {
	ctl := &FileController{
		log: log.WithOptions(zap.Fields(zap.String("module", "FileController"))),
	}
	v1 := apiV1.Group
	file := v1.Group("/file")
	{
		file.GET("/list", ctl.ListFile)
		file.GET("/upload", ctl.UploadLocal)
		file.GET("/oss", ctl.UploadOss)
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
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
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
