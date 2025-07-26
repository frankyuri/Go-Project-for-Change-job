package models

import "time"

// Category 商品分類資料模型
// PK: ID (主鍵)
// FK: ParentID -> Category.ID (自關聯，支援多層分類)
type Category struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`    // PK: 分類唯一識別碼
	Name      string     `gorm:"size:100;not null" json:"name"`         // 分類名稱
	ParentID  *uint      `gorm:"index" json:"parent_id"`                // FK: 父分類ID，頂層分類為 nil
	Parent    *Category  `gorm:"foreignKey:ParentID" json:"parent"`     // 關聯查詢：自關聯到父分類
	Children  []Category `gorm:"foreignKey:ParentID" json:"children"`   // 關聯查詢：一對多關聯到子分類
	Products  []Product  `gorm:"foreignKey:CategoryID" json:"products"` // 關聯查詢：一對多關聯到商品
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`      // 建立時間
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`      // 更新時間
}
