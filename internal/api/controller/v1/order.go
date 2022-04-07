package v1

import "github.com/gin-gonic/gin"

type OrderController struct {
}

func NewOrderController(engine *gin.Engine) {
	ctl := &OrderController{}
	user := engine.Group("/api/v1")
	{
		user.POST("/order", ctl.Order)
	}
}

func (u OrderController) Order(ctx *gin.Context) {
	ctx.JSON(200, map[string]interface{}{
		"message": "order",
	})
}
