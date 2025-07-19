package main

import (
	"context"
	"go-train/database"
	"go-train/middleware"
	"go-train/routes"
	"go-train/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func main() {
	// 讀取env
	if err := godotenv.Load(); err != nil {
		log.Println("找不到 .env，將使用系統環境變數")
	}
	ngrokToken := os.Getenv("NGROK_AUTHTOKEN")
	lineSecret := os.Getenv("LINE_CHANNEL_SECRET")
	lineToken := os.Getenv("LINE_MESSAGE_API_TOKEN")
	// 初始化數據庫
	database.ConnectDB()

	// 初始化 Gin
	r := gin.Default()
	r.Use(middleware.Logger())
	routes.SetupRoutes(r)

	// 建立 context
	ctx := context.Background()
	tun, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(config.WithDomain("")), // 留空自動分配免費 ngrok domain
		ngrok.WithAuthtoken(ngrokToken),
	)
	if err != nil {
		log.Fatal(err)
	}
	// 印出 ngrok 公網網址
	log.Println("ngrok public url:", tun.URL())
	utils.InitLineBot(lineSecret, lineToken)
	// 啟動伺服器
	// r.Run(":303")
	if err := r.RunListener(tun); err != nil {
		log.Fatal(err)
	}

}
