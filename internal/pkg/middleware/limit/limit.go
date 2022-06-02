package middleware

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

const (
	_rate    = 3
	_cap     = 10
	_timeout = 500
)

func Limiter() gin.HandlerFunc {
	limiters := &sync.Map{}
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		l, _ := limiters.LoadOrStore(ip, rate.NewLimiter(_rate, _cap))
		c, cancel := context.WithTimeout(ctx, time.Millisecond*_timeout)
		defer cancel()
		if err := l.(*rate.Limiter).Wait(c); err != nil {
			ctx.JSON(http.StatusTooManyRequests, err.Error())
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
