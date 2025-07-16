package models

import "time"

type OperationLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Action    string    `json:"action"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
}
