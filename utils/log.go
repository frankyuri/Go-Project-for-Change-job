package utils

import (
	"go-train/models"
	"time"

	"go-train/repositories"

	"github.com/gin-gonic/gin"
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
	repositories.CreateOperationLog(&log)
}

// 在 utils/log.go 中新增
func WriteOperationLogFromContext(
	c *gin.Context,
	operationType string,
	objectID string,
	beforeContent string,
	afterContent string,
	result string,
	module string,
	description string,
	duration int64,
) {
	// 從 JWT 中取得用戶資訊
	// userID, _ := c.Get("userID")
	username, _ := c.Get("username")

	WriteOperationLog(
		username.(string), // operator
		operationType,
		objectID,
		beforeContent,
		afterContent,
		result,
		c.ClientIP(), // ipAddress
		"API",        // source
		module,
		description,
		duration,
	)
}
