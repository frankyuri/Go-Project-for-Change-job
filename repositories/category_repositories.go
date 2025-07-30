package repositories

import (
	"go-train/database"
	"go-train/models"
)

func CreateCategory(category *models.Category) error {
	return database.DB.Create(category).Error
}

func GetCategories(categories *[]models.Category) error {
	return database.DB.Preload("Children").Find(categories).Error
}

func GetCategoryByID(id string, category *models.Category) error {
	return database.DB.Preload("Children").First(category, id).Error
}

func UpdateCategory(category *models.Category) error {
	return database.DB.Save(category).Error
}

func DeleteCategory(id string) error {
	return database.DB.Delete(&models.Category{}, id).Error
}
