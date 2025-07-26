package models

import "time"

// Product 商品資料模型
// PK: ID (主鍵)
// FK: CategoryID -> Category.ID (外鍵關聯到分類表)
type Product struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`       // PK: 商品唯一識別碼
	Name        string    `gorm:"size:255;not null" json:"name"`            // 商品名稱
	Description string    `gorm:"type:text" json:"description"`             // 商品描述
	Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"` // 商品價格
	Stock       int       `gorm:"not null;default:0" json:"stock"`          // 庫存數量
	CategoryID  uint      `gorm:"not null" json:"category_id"`              // FK: 關聯到 Category.ID
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`    // 關聯查詢：一對一關聯到分類
	ImageURL    string    `gorm:"size:500" json:"image_url"`                // 商品圖片連結
	Status      string    `gorm:"size:20;default:'active'" json:"status"`   // 商品狀態 (active/inactive)
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`         // 建立時間
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`         // 更新時間
}
