package v1

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/admin/version"
	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/internal/pkg/middleware/auth"
	"github.com/GodYao1995/Goooooo/internal/pkg/middleware/permission"
	"github.com/GodYao1995/Goooooo/internal/pkg/session"
	"github.com/GodYao1995/Goooooo/pkg/tools"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileController struct {
	log    *zap.Logger
	file   string
	upload string
	store  *session.RedisStore
	logic  domain.FileLogicFace
	perm   *casbin.Enforcer
}

func NewFileController(apiV1 *version.APIV1, log *zap.Logger, store *session.RedisStore, logic domain.FileLogicFace, perm *casbin.Enforcer) {
	ctl := &FileController{
		log:    log.WithOptions(zap.Fields(zap.String("module", "FileController"))),
		file:   "file",
		upload: "../upload/",
		store:  store,
		logic:  logic,
		perm:   perm,
	}
	v1 := apiV1.Group
	file := v1.Group("/file").Use(auth.AuthMiddleware(store))
	{
		file.GET("/list", ctl.ListFile)
		file.POST("/upload", ctl.UploadLocal)
		file.POST("/oss", ctl.UploadOss)
	}
	needAuth := v1.Group("/file").Use(permission.PermissionMiddleware(perm))
	{
		needAuth.GET("/download", ctl.DownloadLocal)
		needAuth.GET("/stream", ctl.DownloadLocalFileStream)
	}
}

// ListFile
// @Summary List File
// @Description List File
// @Tags File
// @Accept  json
// @Produce json
// @Router /list [GET]
func (f FileController) ListFile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

// UploadLocal
// @Summary UploadLocal
// @Description List File
// @Tags File
// @Accept multipart/form-data
// @Param file formData file true "文件上传"
// @Produce json
// @Router /upload [POST]
func (f FileController) UploadLocal(ctx *gin.Context) {
	resp := types.CommonResponse{Code: 1}
	fileObj, err := ctx.FormFile(f.file)
	if err != nil {
		resp.Message = errno.ErrorUploadFile.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	target := f.upload + strconv.Itoa(int(tools.SnowId())) + fileObj.Filename
	if err = ctx.SaveUploadedFile(fileObj, target); err != nil {
		resp.Message = errno.ErrorUploadFile.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	var user uint64
	if v, ok := ctx.Get("user"); ok {
		session := v.(domain.UserSession)
		user = session.Id
	} else {
		resp.Message = errno.ErrorUserNotLogin.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	// Just for testing purposes 1526448643605794816 [Temp]
	err = f.logic.UploadFile(fileObj.Filename, fileObj.Size, target, 1526448643605794816, user)
	if err != nil {
		resp.Message = err.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp.Code = 0
	resp.Message = errno.UploadSuccess
	ctx.JSON(http.StatusOK, resp)
}

// DownloadLocal
// @Summary Download Local File
// @Description Download Local File
// @Tags File
// @Accept json
// @Param q query string true  "download file"
// @Produce json
// @Success 0 {object} application/octet-stream
// @Failure 1 {object}
// @Router /download [GET]
func (f FileController) DownloadLocal(ctx *gin.Context) {
	path := f.upload
	filename := ctx.Query("filename")
	ctx.FileAttachment(path+filename, filename)
}

// DownloadLocalFileStream
// @Summary DownloadLocalFileStream
// @Description DownloadLocalFileStream
// @Tags File
// @Accept json
// @Param q query string true  "download file"
// @Produce json
// @Success 0 {object} application/octet-stream
// @Failure 1 {object}
// @Router /stream [GET]
func (f FileController) DownloadLocalFileStream(ctx *gin.Context) {
	resp := types.CommonResponse{Code: 1}
	path := f.upload
	filename := ctx.Query("filename")
	sourceFile, err := os.Open(path + filename)
	if err != nil {
		resp.Message = err.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	data, err := io.ReadAll(sourceFile)
	if err != nil {
		resp.Message = err.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	ctx.Data(http.StatusOK, "application/octet-stream", data)
}

// UploadOss
// @Summary UploadOss
// @Description Upload Oss
// @Tags File
// @Accept multipart/form-data
// @Param file formData file true "文件上传Oss"
// @Produce json
// @Router /oss [POST]
func (f FileController) UploadOss(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
