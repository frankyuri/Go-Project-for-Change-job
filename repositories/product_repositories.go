package repositories

import (
	"go-train/database"
	"go-train/models"
)

func CreateProduct(product *models.Product) error {
	return database.DB.Create(product).Error
}
