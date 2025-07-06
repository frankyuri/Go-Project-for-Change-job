// repositories/user_repository.go
package repositories

import (
	"fmt"
	"go-train/database"
	"go-train/models"

	"gorm.io/gorm/clause"
)

func CreateUser(user *models.User) error {
	database.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoNothing: true,
	})
	result := database.DB.Create(user)
	return result.Error
}

func GetUserByID(id int) (models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	return user, result.Error
}

// func GetUsers() ([]models.User, error) {
// 	var users []models.User
// 	result := database.DB.Find(&users)
// 	return users, result.Error
// }

func DeleteUser(id int) error {
	result := database.DB.Delete(&models.User{}, id)
	return result.Error
}

func GetUsers(page, pageSize int, sortBy, sortOrder string) ([]models.User, int64, error) {
	var users []models.User
	var count int64

	// 獲取總數
	if err := database.DB.Model(&models.User{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// 計算偏移量
	offset := (page - 1) * pageSize

	// 排序方向
	direction := "asc"
	if sortOrder == "desc" {
		direction = "desc"
	}

	// 構建排序查詢
	orderQuery := fmt.Sprintf("%s %s", sortBy, direction)

	// 執行分頁查詢
	result := database.DB.Order(orderQuery).Offset(offset).Limit(pageSize).Find(&users)

	return users, count, result.Error
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
