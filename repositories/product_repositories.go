package repositories

import (
	"go-train/database"
	"go-train/models"
)

func CreateProduct(product *models.Product) error {
	return database.DB.Create(product).Error
}

func GetProducts(products *[]models.Product) error {
	return database.DB.Preload("Category").Find(products).Error
}

func GetProductByID(id string, product *models.Product) error {
	return database.DB.Preload("Category").First(product, id).Error
}

func UpdateProduct(product *models.Product) error {
	return database.DB.Save(product).Error
}

func DeleteProduct(id string) error {
	return database.DB.Delete(&models.Product{}, id).Error
}
