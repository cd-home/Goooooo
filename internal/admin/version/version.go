package version

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewAPIV1, NewAPIV2)

type APIV1 struct {
	Group   *gin.RouterGroup
	Version string
}

func NewAPIV1(engine *gin.Engine) *APIV1 {
	version := "/api/v1"
	return &APIV1{Group: engine.Group(version), Version: version}
}

type APIV2 struct {
	Group   *gin.RouterGroup
	Version string
}

func NewAPIV2(engine *gin.Engine) *APIV2 {
	version := "/api/v1"
	return &APIV2{Group: engine.Group(version), Version: version}
}
