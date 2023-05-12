package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] || [%s %s\t%s - %s] %s %d %s (%s)\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.MethodColor(),
			params.Method,
			params.ResetColor(),
			params.Path,
			params.StatusCodeColor(),
			params.StatusCode,
			params.ResetColor(),
			params.Latency,
		)
	})
}
