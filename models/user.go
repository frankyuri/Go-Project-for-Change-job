package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	DisplayID uint   `gorm:"autoIncrement"`
	Username  string `json:"username" gorm:"uniqueIndex;not null"`
	Password  string `json:"password" gorm:"not null"`
	Email     string `json:"email" gorm:"uniqueIndex;not null"`
}

type Counter struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"uniqueIndex"`
	Value uint
}

// 獲取下一個用戶DisplayID
func GetNextUserDisplayID(db *gorm.DB) (uint, error) {
	var counter Counter

	// 查找或創建用戶計數器
	err := db.Where("name = ?", "user_display_id").First(&counter).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 創建新的計數器
			counter = Counter{Name: "user_display_id", Value: 1}
			if err := db.Create(&counter).Error; err != nil {
				return 0, err
			}
			return 1, nil
		}
		return 0, err
	}

	// 更新計數器
	counter.Value++
	if err := db.Save(&counter).Error; err != nil {
		return 0, err
	}

	return counter.Value, nil
}
