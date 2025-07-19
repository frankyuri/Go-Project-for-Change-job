package utils

import (
	"go-train/database"
	"go-train/models"
	"time"
)

func WriteOperationLog(
	operator string,
	operationType string,
	objectID string,
	beforeContent string,
	afterContent string,
	result string,
	ipAddress string,
	source string,
	module string,
	description string,
	duration int64,
) {
	log := models.OperationLog{
		Operator:      operator,
		OperationType: operationType,
		ObjectID:      objectID,
		BeforeContent: beforeContent,
		AfterContent:  afterContent,
		Result:        result,
		IPAddress:     ipAddress,
		Source:        source,
		Module:        module,
		Description:   description,
		Duration:      duration,
		CreatedAt:     time.Now(),
	}
	database.DB.Create(&log)
}
