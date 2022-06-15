package v1

import (
	"context"

	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/admin/version"
	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/consts"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/internal/pkg/middleware/auth"
	"github.com/GodYao1995/Goooooo/internal/pkg/middleware/tracer"
	"github.com/GodYao1995/Goooooo/internal/pkg/res"
	"github.com/GodYao1995/Goooooo/internal/pkg/session"
	"github.com/GodYao1995/Goooooo/pkg/xhttp/param"
	"github.com/GodYao1995/Goooooo/pkg/xtracer"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type UserController struct {
	logic domain.UserLogicFace
	log   *zap.Logger
	store *session.RedisStore
}

func NewUserController(
	apiV1 *version.APIV1,
	log *zap.Logger,
	logic domain.UserLogicFace,
	store *session.RedisStore, xtracer *xtracer.XTracer) {
	ctl := &UserController{
		logic: logic,
		log:   log.WithOptions(zap.Fields(zap.String("module", "UserController"))),
		store: store,
	}

	// API version
	v1 := apiV1.Group.Group("/user").Use(tracer.Tracing(xtracer))

	// No Need Authorization
	{
		v1.POST("/register", ctl.Register)
		v1.POST("/login", ctl.Login)
	}

	// Need Authorization
	needAuth := v1.Use(auth.AuthMiddleware(store))
	{
		needAuth.GET("/profile", ctl.GetUserProfile)
		needAuth.GET("/retrieves", ctl.GetAllUser)
		needAuth.POST("/logout", ctl.Logout)
		needAuth.POST("/modify_password", ctl.ModifyPassword)
	}
}

// Login
// @Summary User Login
// @Description User Login
// @Tags User
// @Accept  json
// @Produce json
// @Param login body types.LoginParam true "login"
// @Success 0 {object} domain.UserVO {"code":0,"data": domain.UserVO, "msg":"Success"}
// @Failure 1 {object} types.CommonResponse {"code":1,"data":null,"msg":"Error"}
// @Router /login [POST]
func (u UserController) Login(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "UserController-Login")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("Controller", "Login")
		span.Finish()
	}()
	params := types.LoginParam{}
	resp := res.CommonResponse{Code: 1}
	if ok, valid := param.ShouldBindJSON(ctx, &params); !ok {
		resp.Message = valid
		resp.Failure(ctx)
		return
	}
	view, err := u.logic.Login(next, ctx.Request, ctx.Writer, params.Account, params.Password)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = view
		resp.Code = 0
		resp.Message = errno.LoginSuccess
	}
	resp.Success(ctx)
}

// Register
// @Summary User Register
// @Description User Register
// @Tags User
// @Accept  json
// @Produce json
// @Param register body types.RegisterParam true "register"
// @Success 0 {object} types.CommonResponse {"code":1,"data":null,"msg":"Success"}
// @Failure 1 {object} types.CommonResponse {"code":0,"data":null,"msg":"Error"}
// @Router /register [POST]
func (user UserController) Register(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "UserController-Register")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("Controller", "Register")
		span.Finish()
	}()
	params := types.RegisterParam{}
	resp := res.CommonResponse{Code: 1}
	if ok, valid := param.ShouldBindJSON(ctx, &params); !ok {
		resp.Message = valid
		resp.Failure(ctx)
		return
	}
	if err := user.logic.Register(next, params.Account, params.Password); err != nil {
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = errno.RegisterSuccess
	}
	resp.Success(ctx)
}

// GetUserProfile
// @Summary Get UserProfile
// @Description Get UserProfile
// @Tags User
// @Accept  json
// @Produce json
// @Router /profile [GET]
func (u UserController) GetUserProfile(ctx *gin.Context) {
	if v, ok := ctx.Get(consts.SROREKEY); ok {
		session := v.(domain.UserSession)
		ctx.JSON(200, map[string]interface{}{
			"message": session,
		})
	}
}

// GetAllUser
// @Summary Get All User
// @Description Get All User
// @Tags User
// @Accept  json
// @Produce json
// @Router /users [GET]
func (u UserController) GetAllUser(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "UserController-GetAllUser")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("Controller", "GetAllUser")
		span.Finish()
	}()
	resp := res.CommonResponse{Code: 1}
	view, err := u.logic.RetrieveAllUser(next)
	if err != nil {
		resp.Message = err.Error()
		resp.Failure(ctx)
	} else {
		resp.Data = view
		resp.Code = 0
		resp.Message = errno.Success
	}
	resp.Success(ctx)
}

// GetUserProfile
// @Summary Get UserProfile
// @Description Get UserProfile
// @Tags User
// @Accept  json
// @Produce json
// @Router /profile [GET]
func (u UserController) Logout(ctx *gin.Context) {
	resp := res.CommonResponse{Code: 1}
	err := u.logic.Logout(ctx, ctx.Request, ctx.Writer)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = errno.LogOutSuccess
	}
	resp.Success(ctx)
}

// ModifyPassword
// @Summary Modify Password
// @Description Modify Password
// @Tags User
// @Accept  json
// @Produce json
// @Router /modify_password [GET]
func (u UserController) ModifyPassword(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "UserController-ModifyPassword")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("UserController", "ModifyPassword")
		span.Finish()
	}()
	params := types.ModifyPasswordParam{}
	resp := res.CommonResponse{Code: 1}
	if ok, valid := param.ShouldBindJSON(ctx, &params); !ok {
		resp.Message = valid
		span.LogKV("bind err", valid)
		resp.Failure(ctx)
		return
	}
	// Current user
	var uid uint64
	if v, ok := ctx.Get(consts.SROREKEY); ok {
		session := v.(domain.UserSession)
		uid = session.Id
	} else {
		resp.Message = errno.ErrorSessionsInvalid.Error()
		span.LogKV("session", errno.ErrorSessionsInvalid.Error())
		resp.Failure(ctx)
		return
	}
	// Modify Password
	if err := u.logic.ModifyPassword(next, params.OriginPassword, params.NewPassword, uid); err != nil {
		resp.Message = err.Error()
		resp.Failure(ctx)
		return
	}
	// Logout
	if err := u.logic.Logout(next, ctx.Request, ctx.Writer); err != nil {
		resp.Message = errno.ErrorLogOutForced.Error()
	} else {
		resp.Code = 0
		resp.Message = errno.ModifyPasswordSuccess
	}
	resp.Success(ctx)
}
