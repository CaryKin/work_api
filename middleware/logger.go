package middleware

import (
	"gin-vue-admin/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// ZapLogger 接收gin框架的路由日志
func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		cost := time.Since(start)

		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if errorMessage == "" {
			if global.GVA_CONFIG.Debug {
				global.GVA_LOG.Info(path,
					zap.Int("status", c.Writer.Status()),
					zap.String("method", c.Request.Method),
					zap.String("path", path),
					zap.String("query", query),
					zap.String("ip", c.ClientIP()),
					zap.String("user-agent", c.Request.UserAgent()),
					zap.Duration("cost", cost),
				)
			}
		} else {
			global.GVA_LOG.Error(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("cost", cost),
			)
		}
	}
}
