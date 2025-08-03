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
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Println("找不到 .env，將使用系統環境變數")
	}
	ngrokToken := os.Getenv("NGROK_AUTHTOKEN")
	lineSecret := os.Getenv("LINE_CHANNEL_SECRET")
	lineToken := os.Getenv("LINE_MESSAGE_API_TOKEN")
	// 初始化數據庫
	database.ConnectDB()

	// 初始化 Gin
	r := gin.Default()
	utils.InitRedis()
	// 信任代理，讓 Gin 正確處理 X-Forwarded-* headers
	r.SetTrustedProxies([]string{"0.0.0.0/0"})

	// 加入 CORS 中間件
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	r.Use(middleware.Logger())
	routes.SetupRoutes(r)

	// 建立 context
	ctx := context.Background()
	tun, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(), // 使用預設 HTTP endpoint，ngrok 會自動處理 HTTPS
		ngrok.WithAuthtoken(ngrokToken),
	)
	if err != nil {
		log.Fatal(err)
	}
	// 印出 ngrok 公網網址
	log.Println("ngrok public url:", tun.URL())
	utils.InitLineBot(lineSecret, lineToken)
	// 啟動伺服器
	// r.Run(":3030")
	if err := r.RunListener(tun); err != nil {
		log.Fatal(err)
	}
}
