package middleware

import (
	"auth-app/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(ctx *gin.Context) {
	log := logger.Log
	startTime := time.Now()

	ctx.Next()

	endTime := time.Now()
	latency := endTime.Sub(startTime).String()
	reqMethod := ctx.Request.Method
	reqHost := ctx.Request.Host
	reqURI := ctx.Request.RequestURI
	statusCode := ctx.Writer.Status()
	clientIP := ctx.ClientIP()

	fields := map[string]any{
		"method":    reqMethod,
		"uri":       reqURI,
		"status":    statusCode,
		"latency":   latency,
		"client_ip": clientIP,
		"host":      reqHost,
	}
	
	if lastErr := ctx.Errors.Last(); lastErr != nil {
		log.WithFields(fields).Error(ctx.Errors[0])
		return
	}

	log.WithFields(fields).Infof("REQUEST %s %s SUCCESS", reqMethod, reqURI)
}