package models

import "time"

type LineTodo struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    string    `gorm:"index"` // Line UserID
    Content   string
    Status    string    // "todo", "done", "note" 等
    CreatedAt time.Time
    UpdatedAt time.Time
}