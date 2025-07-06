package controllers

import (
	"errors"
	"go-train/models"
	"go-train/repositories"
	"go-train/utils"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 創建新用戶
func CreateUser(c *gin.Context) {
	var user models.User

	// 改進的請求體綁定和驗證
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, "無效的請求數據: "+err.Error()))
		return
	}

	// 基本數據驗證
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, "郵箱不能為空"))
		return
	}

	// 創建用戶
	if err := repositories.CreateUser(&user); err != nil {
		// 檢查特定錯誤類型，例如唯一約束違反
		if strings.Contains(err.Error(), "unique constraint") || strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, utils.ErrorResponse(http.StatusConflict, "該郵箱已被使用"))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, "創建用戶失敗"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(http.StatusOK, "成功創建用戶", user))
}

// 根據ID獲取單個用戶
func GetUserByID(c *gin.Context) {
	idx := c.Param("id")
	id, err := strconv.Atoi(idx)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, "無效的ID格式"))
		return
	}

	user, err := repositories.GetUserByID(id)

	if err != nil {
		// 區分不同錯誤類型
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, utils.ErrorResponse(http.StatusNotFound, "用戶不存在"))
		} else {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, "查詢用戶失敗"))
		}
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "成功獲取用戶", user))
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// 獲取所有用戶
func GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	sortBy := c.DefaultQuery("sort_by", "id")
	sortOrder := c.DefaultQuery("sort_order", "asc")

	var users []models.User
	users, totalCount, err := repositories.GetUsers(page, pageSize, sortBy, sortOrder)

	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(http.StatusInternalServerError, "獲取用戶列表失敗"))
		return
	}

	// 如果沒有用戶，返回空數組而不是錯誤
	if len(users) == 0 {
		c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "沒有找到用戶", []models.User{}))
		return
	}

	// 添加元數據
	response := map[string]interface{}{
		"users": users,
		"meta": map[string]interface{}{
			"total_count": totalCount,
			"page":        page,
			"page_size":   pageSize,
			"total_pages": int(math.Ceil(float64(totalCount) / float64(pageSize))),
		},
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "成功獲取所有用戶", response))

}

// // UpdateUser 更新用戶
// func UpdateUser(c *gin.Context) {
// 	var user models.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := repositories.UpdateUser(&user); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "成功更新用戶", user))
// }

// // DeleteUser 刪除用戶
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, "無效的ID格式"))
		return
	}

	_, err = repositories.GetUserByID(idx)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(http.StatusNotFound, "用戶不存在"))
		return
	}

	if err := repositories.DeleteUser(idx); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, "無法刪除用戶"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "成功刪除用戶", nil))
}
