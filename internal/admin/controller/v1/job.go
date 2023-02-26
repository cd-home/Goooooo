package v1

import (
	"context"

	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/cd-home/Goooooo/internal/admin/version"
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
	job := v1.Group("/jobs")
	{
		job.GET("/", ctl.GetJobs)
		job.GET("/users2es", ctl.User2EsJobs)
	}
}

// GetJobs
// @Summary Get Jobs
// @Description Get Jobs
// @Tags Job
// @Accept  json
// @Produce json
// @Router /jobs [GET]
func (u JobController) GetJobs(ctx *gin.Context) {
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

// GetJobs
// @Summary Get Jobs
// @Description Get Jobs
// @Tags Job
// @Accept  json
// @Produce json
// @Router /jobs [GET]
func (u JobController) User2EsJobs(ctx *gin.Context) {
	sum := tasks.Signature{
		Name: "user2es",
		Args: []tasks.Arg{},
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
