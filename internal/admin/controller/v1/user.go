package controller

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

func (u UserController) Login(ctx *gin.Context) {
	ctx.JSON(200, map[string]interface{}{
		"message": "ok",
	})
}
