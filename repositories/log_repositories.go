package repositories

import (
	"go-train/database"
	"go-train/models"
)

func CreateOperationLog(log *models.OperationLog) error {
	return database.DB.Create(log).Error
}
