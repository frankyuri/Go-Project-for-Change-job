package controllers

import (
	"fmt"
	"go-train/models"
	"go-train/repositories"
	"go-train/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category models.Category

	start := time.Now()

	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "未授權",
		})
	}

	_, exists = c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse(http.StatusUnauthorized, "無法取得用戶資訊"))
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, "無效的請求數據: "+err.Error()))
	}

	if err := repositories.CreateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, "創建分類失敗: "+err.Error()))
		utils.WriteOperationLogFromContext(c, "create_category", fmt.Sprintf("%d", category.ID), "", "", "fail", "API", "Category", time.Since(start).Milliseconds())
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(http.StatusCreated, "分類創建成功", category))
	utils.WriteOperationLogFromContext(c, "create_category", fmt.Sprintf("%d", category.ID), "", "", "success", "API", "Category", time.Since(start).Milliseconds())
}
