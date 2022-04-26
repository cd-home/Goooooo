package v1

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController(engine *gin.Engine) {
	ctl := &UserController{}
	user := engine.Group("/api/v1")
	{
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
// @Router /register [POST]
func (u UserController) Register(ctx *gin.Context) {
	ctx.JSON(200, map[string]interface{}{
		"message": "ok",
	})
}