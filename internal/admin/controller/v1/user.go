package v1

import (
	"encoding/json"
	"net/http"

	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/admin/version"
	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/internal/pkg/middleware/auth"
	"github.com/GodYao1995/Goooooo/internal/pkg/session"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
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
	v1 := apiV1.Group
	// No Need Authorization
	user := v1.Group("/user")
	{
		user.POST("/register", ctl.Register)
		user.POST("/login", ctl.Login)
	}
	// Need Authorization
	needAuth := v1.Group("/user")
	needAuth.Use(auth.AuthMiddleware(store))
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
// @Router /login [POST]
func (u UserController) Login(ctx *gin.Context) {
	params := types.LoginParam{}
	resp := types.CommonResponse{Code: 0}
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
	session, _ := u.store.Get(ctx.Request, "SESSIONID")
	// store session
	values, _ := json.Marshal(obj)
	session.Values["user"] = values
	// TODO 后期修改到配置项
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 2,
		HttpOnly: true,
		Secure:   true,
	}
	// write Cookie and store to Redis Session
	if err := session.Save(ctx.Request, ctx.Writer); err != nil {
		resp.Message = err.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = view
	resp.Message = errno.Success
	ctx.JSON(http.StatusOK, resp)
}

// Register
// @Summary User Register
// @Description User Register
// @Tags User
// @Accept  json
// @Produce json
// @Param register body types.RegisterParam true "register"
// @Success 1 {object} types.CommonResponse {"code":1,"data":null,"msg":"Success"}
// @Failure 0 {object} types.CommonResponse {"code":0,"data":null,"msg":"Error"}
// @Router /register [POST]
func (user UserController) Register(ctx *gin.Context) {
	params := types.RegisterParam{}
	resp := types.CommonResponse{Code: 0}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		resp.Message = errno.ErrorParamsParse.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	if err := user.logic.Register(ctx, params.Account, params.Password); err != nil {
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = errno.Success
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
