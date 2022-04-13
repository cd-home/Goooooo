package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// NoCache prevent the client from caching the HTTP response
func NoCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
		c.Header("Expires", "Thu, 01 Jan 1970 2020_01:2020_01:2020_01 GMT")
		c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
		c.Next()
	}
}

// Cors for options requests and aborts then exits the middleware chain and ends the request
// 解决跨域的是先发一次options请求，获取AllowHeader，允许跨域之后才会再发真正的Post请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// Secure is a middleware function that appends security and resource access headers.
func Secure() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")
		if c.Request.TLS != nil {
			c.Header("Strict-Transport-Security", "max-age=31536000")
		}
	}
}

// RequestID is a middleware function that appends RequestID.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		_requestId := uuid.NewV4().String()
		c.Set("X-Request-Id", _requestId)
		c.Writer.Header().Set("X-Request-Id", _requestId)
		c.Next()
	}
}
