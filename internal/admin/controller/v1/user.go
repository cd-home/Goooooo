package v1

import (
	"encoding/json"
	"net/http"

	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/internal/pkg/session"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	logic domain.UserLogicFace
	log   *zap.Logger
	store *session.RedisStore
}

func NewUserController(engine *gin.Engine, log *zap.Logger, logic domain.UserLogicFace, store *session.RedisStore) {
	ctl := &UserController{
		logic: logic,
		log:   log.WithOptions(zap.Fields(zap.String("module", "UserController"))),
		store: store,
	}
	user := engine.Group("/api/v1/user")
	{
		user.POST("/register", ctl.Register)
		user.POST("/login", ctl.Login)
		user.GET("/profile", ctl.GetUserProfile)
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
		resp.Message = errno.ParamsParseError.Error()
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
	// store
	sessions, _ := json.Marshal(obj)
	session.Values["user"] = sessions
	if err := u.store.Save(ctx.Request, ctx.Writer, session); err != nil {
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
		resp.Message = errno.ParamsParseError.Error()
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
	session, _ := u.store.New(ctx.Request, "SESSIONID")
	var user domain.UserSession
	json.Unmarshal(session.Values["user"].([]byte), &user)
	ctx.JSON(200, map[string]interface{}{
		"message": user,
	})
}
