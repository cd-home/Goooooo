package types

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CommonResponse [Response]
type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success
func (c CommonResponse) Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c)
}

// Failure
func (c CommonResponse) Failure(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c)
}
