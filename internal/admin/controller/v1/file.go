package v1

import (
	"context"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/admin/version"
	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/consts"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/internal/pkg/middleware/auth"
	"github.com/GodYao1995/Goooooo/internal/pkg/middleware/permission"
	"github.com/GodYao1995/Goooooo/internal/pkg/res"
	"github.com/GodYao1995/Goooooo/internal/pkg/session"
	"github.com/GodYao1995/Goooooo/pkg/tools"
	"github.com/GodYao1995/Goooooo/pkg/xhttp/param"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type FileController struct {
	log    *zap.Logger
	file   string
	upload string
	logic  domain.FileLogicFace
}

func NewFileController(apiV1 *version.APIV1, log *zap.Logger, store *session.RedisStore, logic domain.FileLogicFace, perm *casbin.Enforcer) {
	ctl := &FileController{
		log:    log.WithOptions(zap.Fields(zap.String("module", "FileController"))),
		file:   "file",
		upload: "../upload/",
		logic:  logic,
	}

	// API version
	v1 := apiV1.Group.Group("/file")

	// Need Authorization
	needAuth := v1.Use(auth.AuthMiddleware(store))
	{
		needAuth.GET("/list", ctl.ListFile)
		needAuth.POST("/upload", ctl.UploadLocal)
		needAuth.POST("/oss", ctl.UploadOss)
	}

	// Need Authorization And Permission
	needPerm := needAuth.Use(permission.PermissionMiddleware(perm))
	{
		needPerm.GET("/download", ctl.DownloadLocal)
		needPerm.GET("/stream", ctl.DownloadLocalFileStream)
		needPerm.DELETE("/delete", ctl.DeleteLocal)
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
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "FileController-UploadLocal")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("FileController", "UploadLocal")
		span.Finish()
	}()
	resp := res.CommonResponse{Code: 1}
	fileObj, err := ctx.FormFile(f.file)
	if err != nil {
		resp.Message = errno.ErrorUploadFile.Error()
		resp.Failure(ctx)
		return
	}
	target := f.upload + strconv.Itoa(int(tools.SnowId())) + fileObj.Filename
	if err = ctx.SaveUploadedFile(fileObj, target); err != nil {
		resp.Message = errno.ErrorUploadFile.Error()
		resp.Failure(ctx)
		return
	}
	var user uint64
	if v, ok := ctx.Get(consts.SROREKEY); ok {
		session := v.(domain.UserSession)
		user = session.Id
	} else {
		resp.Message = errno.ErrorUserNotLogin.Error()
		resp.Failure(ctx)
		return
	}
	// Just for testing purposes 1526448643605794816 [Temp]
	err = f.logic.UploadFile(next, fileObj.Filename, fileObj.Size, target, 1526448643605794816, user)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = errno.UploadFileSuccess
	}
	resp.Success(ctx)
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

// DeleteLocal
// @Summary DeleteLocal File
// @Description DeleteLocal File
// @Tags File
// @Accept json
// @Produce json
// @Param file_id query
// @Router /delete [DELETE]
func (f FileController) DeleteLocal(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "FileController-DeleteLocal")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("FileController", "DeleteLocal")
		span.Finish()
	}()
	params := types.DeleteFileParam{}
	resp := res.CommonResponse{Code: 1}
	if ok, valid := param.ShouldBindQuery(ctx, &params); !ok {
		resp.Message = valid
		resp.Failure(ctx)
		return
	}
	err := f.logic.DeleteFile(next, params.FileId)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = errno.Success
	}
	resp.Success(ctx)
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
	resp := res.CommonResponse{Code: 1}
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
