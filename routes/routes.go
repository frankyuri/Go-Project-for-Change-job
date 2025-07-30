package routes

import (
	"go-train/controllers"
	"go-train/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	//webhook 路由
	r.POST("/line/webhook", controllers.LineWebhook)
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
		api.POST("/products", controllers.CreateProduct)
		api.POST("/categories", controllers.CreateCategory)
		api.GET("/products", controllers.GetProducts)
		api.GET("/products/:id", controllers.GetProductByID)
		api.PUT("/products/:id", controllers.UpdateProduct)
		api.DELETE("/products/:id", controllers.DeleteProduct)
		api.GET("/categories", controllers.GetCategories)
		api.GET("/categories/:id", controllers.GetCategoryByID)
		api.PUT("/categories/:id", controllers.UpdateCategory)
		api.DELETE("/categories/:id", controllers.DeleteCategory)
	}
}
