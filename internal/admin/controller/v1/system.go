package v1

import (
	"github.com/GodYao1995/Goooooo/internal/admin/version"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type SysController struct {
	db *sqlx.DB
}

func NewSysController(apiV1 *version.APIV1, db *sqlx.DB) {
	v1 := apiV1.Group
	ctl := &SysController{db: db}
	health := v1.Group("/health")
	{
		health.GET("/sys", ctl.SysCheck)
		health.GET("/db", ctl.DbCheck)
	}
}

// SysCheck
// @Summary Sys HealthyCheck
// @Description Sys HealthyCheck
// @Tags Sys
// @Accept  json
// @Produce json
// @Router /health [GET]
func (u SysController) SysCheck(ctx *gin.Context) {
	ctx.JSON(200, map[string]interface{}{
		"message": "ok",
	})
}

// DBStats
// @Summary Sys DBStats
// @Description Sys DBStats
// @Tags Sys
// @Accept  json
// @Produce json
// @Router /db [GET]
func (sys SysController) DbCheck(ctx *gin.Context) {
	stats := sys.db.Stats()
	ctx.JSON(200, map[string]interface{}{
		"message": "ok",
		"stats":   stats,
	})
}
