package middleware

import (
	"fmt"
	"go-train/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.ErrorResponse(http.StatusUnauthorized, "未提供認證令牌"))
			c.Abort()
			return
		}

		// 檢查 Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, utils.ErrorResponse(http.StatusUnauthorized, "無效的認證格式"))
			c.Abort()
			return
		}

		// Debug: 印出 token 前幾個字元
		token := parts[1]
		if len(token) > 10 {
			fmt.Printf("Token: %s...\n", token[:10])
		} else {
			fmt.Printf("Token: %s\n", token)
		}

		claims, err := utils.ValidateToken(token)
		if err != nil {
			fmt.Printf("Token validation error: %v\n", err)
			c.JSON(http.StatusUnauthorized, utils.ErrorResponse(http.StatusUnauthorized, "無效的令牌"))
			c.Abort()
			return
		}

		// 將用戶信息存儲到上下文中
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
