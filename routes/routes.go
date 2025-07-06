package routes

import (
	"go-train/controllers"
	"go-train/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 公開路由
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterUser)
		auth.POST("/login", controllers.LoginUser)
	}

	// 需要認證的路由
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUserByID)
		api.DELETE("/users/:id", controllers.DeleteUser)
		api.POST("/change-password", controllers.ChangePassword)
		// ... 其他需要認證的路由
	}
}
