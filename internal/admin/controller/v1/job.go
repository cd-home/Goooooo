package v1

import (
	"context"

	"github.com/GodYao1995/Goooooo/internal/admin/version"
	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type JobController struct {
	db        *sqlx.DB
	jobServer *machinery.Server
}

func NewJobController(apiV1 *version.APIV1, db *sqlx.DB, jobServer *machinery.Server) {
	v1 := apiV1.Group
	ctl := &JobController{db: db, jobServer: jobServer}
	health := v1.Group("/health")
	{
		health.GET("/job", ctl.Job)
	}
}

// DBStats
// @Summary Sys DBStats
// @Description Sys DBStats
// @Tags Sys
// @Accept  json
// @Produce json
// @Router /db [GET]
func (u JobController) Job(ctx *gin.Context) {
	sum := tasks.Signature{
		Name: "sum",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
	}

	_, err := u.jobServer.SendTaskWithContext(context.Background(), &sum)
	if err != nil {
		ctx.JSON(200, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, map[string]interface{}{
		"message": "ok",
	})
}
