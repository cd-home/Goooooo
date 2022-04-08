package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type SysController struct {
	db *sqlx.DB
}

func NewSysController(engine *gin.Engine, db *sqlx.DB) {
	ctl := &SysController{db: db}
	user := engine.Group("/api/v1")
	{
		user.GET("/db", ctl.DBStats)
	}
}

// DBStats
// @Summary Sys DBStats
// @Description Sys DBStats
// @Tags Sys
// @Accept  json
// @Produce json
// @Router /db [GET]
func (u SysController) DBStats(ctx *gin.Context) {
	stats := u.db.Stats()
	ctx.JSON(200, map[string]interface{}{
		"message": "ok",
		"stats":   stats,
	})
}
