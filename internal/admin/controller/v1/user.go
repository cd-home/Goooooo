package v1

import (
	"net/http"

	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/admin/version"
	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/internal/pkg/middleware/auth"
	"github.com/GodYao1995/Goooooo/internal/pkg/session"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	logic domain.UserLogicFace
	log   *zap.Logger
	store *session.RedisStore
}

func NewUserController(apiV1 *version.APIV1, log *zap.Logger, logic domain.UserLogicFace, store *session.RedisStore) {
	ctl := &UserController{
		logic: logic,
		log:   log.WithOptions(zap.Fields(zap.String("module", "UserController"))),
		store: store,
	}

	// API version
	v1 := apiV1.Group.Group("/user")

	// No Need Authorization
	{
		v1.POST("/register", ctl.Register)
		v1.POST("/login", ctl.Login)
	}
	
	// Need Authorization
	needAuth := v1.Use(auth.AuthMiddleware(store))
	{
		needAuth.GET("/profile", ctl.GetUserProfile)
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
	params := types.LoginParam{}
	resp := types.CommonResponse{Code: 1}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		resp.Message = errno.ErrorParamsParse.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	view, obj, err := u.logic.Login(ctx, params.Account, params.Password)
	if err != nil {
		resp.Message = err.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	// Set Session
	if err := u.logic.SetSession(ctx.Request, ctx.Writer, obj); err != nil {
		resp.Message = err.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = view
	resp.Code = 0
	resp.Message = errno.LoginSuccess
	ctx.JSON(http.StatusOK, resp)
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
	params := types.RegisterParam{}
	resp := types.CommonResponse{Code: 1}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		resp.Message = errno.ErrorParamsParse.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	if err := user.logic.Register(ctx, params.Account, params.Password); err != nil {
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = errno.RegisterSuccess
	}
	ctx.JSON(http.StatusOK, resp)
}

// GetUserProfile
// @Summary Get UserProfile
// @Description Get UserProfile
// @Tags User
// @Accept  json
// @Produce json
// @Router /profile [GET]
func (u UserController) GetUserProfile(ctx *gin.Context) {
	if v, ok := ctx.Get("user"); ok {
		session := v.(domain.UserSession)
		ctx.JSON(200, map[string]interface{}{
			"message": session,
		})
	}
}
