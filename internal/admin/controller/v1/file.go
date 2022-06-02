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
	ctl := &FileController{log: log}
	v1 := apiV1.Group
	file := v1.Group("/file")
	{
		file.GET("/list", ctl.ListFile)
	}
}

// ListFile
// @Summary List File
// @Description List File
// @Tags Directory
// @Accept  json
// @Produce json
// @Router /list [POST]
func (d FileController) ListFile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
