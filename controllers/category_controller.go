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

// 查詢所有分類
func GetCategories(c *gin.Context) {
	var categories []models.Category
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
	if err := repositories.GetCategories(&categories); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, "查詢分類失敗: "+err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "查詢成功", categories))
}

// 查詢單一分類
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
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
	var category models.Category
	if err := repositories.GetCategoryByID(id, &category); err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(http.StatusNotFound, "找不到分類"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "查詢成功", category))
}

// 更新分類
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
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
	var category models.Category
	if err := repositories.GetCategoryByID(id, &category); err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(http.StatusNotFound, "找不到分類"))
		return
	}
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, "無效的請求數據: "+err.Error()))
		return
	}
	category.Name = input.Name
	category.ParentID = input.ParentID
	if err := repositories.UpdateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, "更新分類失敗: "+err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "更新成功", category))
}

// 刪除分類
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
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
	if err := repositories.DeleteCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, "刪除分類失敗: "+err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "刪除成功", nil))
}
