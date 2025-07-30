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

// 建立商品
func CreateProduct(c *gin.Context) {
	var product models.Product
	start := time.Now()

	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授權"})
		return
	}
	_, exists = c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse(http.StatusUnauthorized, "無法取得用戶資訊"))
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, "無效的請求數據: "+err.Error()))
		return
	}
	if err := repositories.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, "創建商品失敗: "+err.Error()))
		utils.WriteOperationLogFromContext(c, "create_product", fmt.Sprintf("%d", product.ID), "", "", "fail", "API", "Product", time.Since(start).Milliseconds())
		return
	}
	c.JSON(http.StatusCreated, utils.SuccessResponse(http.StatusCreated, "商品創建成功", product))
	utils.WriteOperationLogFromContext(c, "create_product", fmt.Sprintf("%d", product.ID), "", "", "success", "API", "Product", time.Since(start).Milliseconds())
}

// 查詢所有商品
func GetProducts(c *gin.Context) {
	var products []models.Product

	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授權"})
		return
	}
	_, exists = c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse(http.StatusUnauthorized, "無法取得用戶資訊"))
		return
	}
	if err := repositories.GetProducts(&products); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, "查詢商品失敗: "+err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "查詢成功", products))
}

// 查詢單一商品
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授權"})
		return
	}
	_, exists = c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse(http.StatusUnauthorized, "無法取得用戶資訊"))
		return
	}
	if err := repositories.GetProductByID(id, &product); err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(http.StatusNotFound, "找不到商品"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "查詢成功", product))
}

// 更新商品
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授權"})
		return
	}
	if err := repositories.GetProductByID(id, &product); err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(http.StatusNotFound, "找不到商品"))
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, "無效的請求數據: "+err.Error()))
		return
	}
	var input models.Product
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock
	product.CategoryID = input.CategoryID
	product.ImageURL = input.ImageURL
	product.Status = input.Status
	if err := repositories.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, "更新商品失敗: "+err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "更新成功", product))
}

// 刪除商品
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授權"})
		return
	}
	_, exists = c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse(http.StatusUnauthorized, "無法取得用戶資訊"))
		return
	}
	if err := repositories.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, "刪除商品失敗: "+err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "刪除成功", nil))
}
