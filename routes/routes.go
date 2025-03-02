package routes

import (
	"go-train/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUserByID)
	r.GET("/ping", controllers.Ping)
	r.GET("/users", controllers.GetUsers)
	r.DELETE("/users/:id", controllers.DeleteUser)
}
