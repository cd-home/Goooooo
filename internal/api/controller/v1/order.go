package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type OrderController struct {
	db *sqlx.DB
}

func NewOrderController(engine *gin.Engine, db *sqlx.DB) {
	ctl := &OrderController{db: db}
	user := engine.Group("/api/v1")
	{
		user.POST("/order", ctl.Order)
	}
}

// Order
// @Summary Goods Order
// @Description Goods Order
// @Tags Order
// @Accept  json
// @Produce json
// @Router /order [POST]
func (u OrderController) Order(ctx *gin.Context) {
	ctx.JSON(200, map[string]interface{}{
		"message": "order",
	})
}
