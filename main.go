package main

import (
	"context"
	"go-train/database"
	"go-train/middleware"
	"go-train/routes"
	"go-train/utils"
	"log"

	"github.com/gin-gonic/gin"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func main() {
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
		ngrok.WithAuthtoken("305Kpph1xTFLHENlCNP58WvkYsn_yzjjX9M9t2RcnGpB67FU"),
	)
	if err != nil {
		log.Fatal(err)
	}
	// 印出 ngrok 公網網址
	log.Println("ngrok public url:", tun.URL())
	utils.InitLineBot("78e8c80bc89ece953217164cda110af3", "hKxIQtfZO+VsgcbPzJ9JntwuhJYqLdtH40+RnpgNJpQCYdsmYyhr6+dKsq2ukK9Lb7762Fyl/xkFUbxBuQvwYUJ6mn8cYNcUIxi1Njkms2u2SOrSlOE0iKRfoNwSgCiuAykiwaLnf9OqM7R9AiajHAdB04t89/1O/w1cDnyilFU=")
	// 啟動伺服器
	// r.Run(":303")
	if err := r.RunListener(tun); err != nil {
		log.Fatal(err)
	}

}
