package models

import "time"

// Order 訂單主檔資料模型
// PK: ID (主鍵)
// FK: UserID -> User.ID (外鍵關聯到用戶表)
type Order struct {
	ID          uint        `gorm:"primaryKey;autoIncrement" json:"id"`               // PK: 訂單唯一識別碼
	OrderNumber string      `gorm:"size:50;uniqueIndex;not null" json:"order_number"` // 訂單編號（唯一）
	UserID      uint        `gorm:"not null" json:"user_id"`                          // FK: 關聯到 User.ID
	User        User        `gorm:"foreignKey:UserID" json:"user"`                    // 關聯查詢：一對一關聯到用戶
	TotalAmount float64     `gorm:"type:decimal(10,2);not null" json:"total_amount"`  // 訂單總金額
	Status      string      `gorm:"size:20;default:'pending'" json:"status"`          // 訂單狀態 (pending/paid/shipped/completed/cancelled)
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`            // 關聯查詢：一對多關聯到訂單明細
	CreatedAt   time.Time   `gorm:"autoCreateTime" json:"created_at"`                 // 建立時間
	UpdatedAt   time.Time   `gorm:"autoUpdateTime" json:"updated_at"`                 // 更新時間
}

// OrderItem 訂單明細資料模型
// PK: ID (主鍵)
// FK: OrderID -> Order.ID (外鍵關聯到訂單主檔)
// FK: ProductID -> Product.ID (外鍵關聯到商品表)
type OrderItem struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`            // PK: 訂單明細唯一識別碼
	OrderID   uint      `gorm:"not null" json:"order_id"`                      // FK: 關聯到 Order.ID
	Order     Order     `gorm:"foreignKey:OrderID" json:"order"`               // 關聯查詢：多對一關聯到訂單
	ProductID uint      `gorm:"not null" json:"product_id"`                    // FK: 關聯到 Product.ID
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`           // 關聯查詢：多對一關聯到商品
	Quantity  int       `gorm:"not null" json:"quantity"`                      // 購買數量
	UnitPrice float64   `gorm:"type:decimal(10,2);not null" json:"unit_price"` // 購買時單價（快照，避免商品價格變動影響）
	SubTotal  float64   `gorm:"type:decimal(10,2);not null" json:"sub_total"`  // 小計金額
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`              // 建立時間
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`              // 更新時間
}
