package repositories

import (
	"go-train/database"
	"go-train/models"
)

func CreateCategory(category *models.Category) error {
	return database.DB.Create(category).Error
}
