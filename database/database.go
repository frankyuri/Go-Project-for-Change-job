// database/database.go
package database

import (
	"fmt"
	"go-train/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }
    
	DB = db
	
	if err:= DB.AutoMigrate(&models.User{}); err != nil {
		panic("Failed to migrate database")
	}


	
	fmt.Println("Database connected!")
}