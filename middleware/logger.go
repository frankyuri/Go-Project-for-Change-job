// middleware/logger.go
package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path
        
        // 處理請求
        c.Next()
        
        // 計算耗時
        latency := time.Since(start)
        log.Printf("API %s - %s - %d - %s", path, c.Request.Method, c.Writer.Status(), latency)
    }
}