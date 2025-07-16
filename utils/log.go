package utils

import (
	"go-train/database"
	"go-train/models"
)

func WriteOperationLog(userID uint, action, detail string) {
	log := models.OperationLog{
		UserID: userID,
		Action: action,
		Detail: detail,
	}
	database.DB.Create(&log)
}
