package v1

import (
	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	logic domain.UserLogicFace
	log   *zap.Logger
}

func NewUserController(engine *gin.Engine, log *zap.Logger, logic domain.UserLogicFace) {
	ctl := &UserController{
		logic: logic,
		log:   log.WithOptions(zap.Fields(zap.String("module", "UserController"))),
	}
	user := engine.Group("/api/v1")
	{
		user.POST("/register", ctl.Register)
		user.POST("/login", ctl.Login)
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
	ctx.JSON(200, map[string]interface{}{
		"message": "ok",
	})
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
	common := types.CommonResponse{}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		common.Code = 0
		common.Message = err.Error()
		ctx.JSON(200, common)
		return
	}
	if err := user.logic.Register(ctx, params.Account, params.Password); err != nil {
		common.Code = 0
		common.Message = err.Error()
		ctx.JSON(200, common)
		return
	} else {
		common.Code = 1
		common.Message = "注册成功"
		ctx.JSON(200, common)
		return
	}
}
