package models

import "time"

type OperationLog struct {
	ID            uint      `gorm:"primaryKey"`
	Operator      string    // 操作者
	OperationType string    // 操作類型
	ObjectID      string    // 目標物件ID
	BeforeContent string    // 操作前內容
	AfterContent  string    // 操作後內容
	Result        string    // 結果
	IPAddress     string    // IP
	Source        string    // 來源
	Module        string    // 模組
	Description   string    // 描述
	Duration      int64     // 執行時間(ms)
	CreatedAt     time.Time // 建立時間
}
