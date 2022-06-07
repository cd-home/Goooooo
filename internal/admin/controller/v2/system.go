package v2

import (
	"github.com/GodYao1995/Goooooo/internal/admin/version"
	"github.com/GodYao1995/Goooooo/internal/pkg/middleware/permission"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type SysController struct {
	db *sqlx.DB
}

func NewSysController(apiV2 *version.APIV2, db *sqlx.DB, e *casbin.Enforcer) {
	ctl := &SysController{db: db}
	v2 := apiV2.Group
	v1 := v2.Group("")
	v1.Use(permission.PermissionMiddleware(e))
	{
		v1.GET("/db", ctl.DBStats)
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
