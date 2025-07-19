package main

import (
	"go-train/database"
	"go-train/middleware"
	"go-train/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 初始化數據庫
	database.ConnectDB()

	// 初始化 Gin
	r := gin.Default()
	r.Use(middleware.Logger())
	routes.SetupRoutes(r)

	// 啟動伺服器
	r.Run(":303")
}
