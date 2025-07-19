package models

import "time"

type TodoStatus string

const (
	StatusTodo TodoStatus = "todo"
	StatusDone TodoStatus = "done"
)

type LineTodo struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    string `gorm:"index"` // Line UserID
	Content   string
	Status    TodoStatus `gorm:"type:varchar(16)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
