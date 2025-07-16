// database/database.go
package database

import (
	"fmt"
	"go-train/models"
	"log"
	"os"

	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	// 加载.env文件（如果使用）
	if err := godotenv.Load(); err != nil {
		panic("Failed to load .env file")
	}
	// 從環境變量讀取配置
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
	// 先連接數據庫
	// 連接數據庫
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	//// 刪除現有的表（如果存在）
	//err = DB.Migrator().DropTable(&models.User{})
	//if err != nil {
	//	log.Fatal("Failed to drop table: ", err)
	//}

	// 創建新的表
	err = DB.AutoMigrate(&models.User{}, &models.Counter{}, &models.OperationLog{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Database migration completed successfully")
}
