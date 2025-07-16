package controllers

import (
	"go-train/database"
	"go-train/models"
	"go-train/repositories"
	"go-train/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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
	var user models.User
	id := c.Param("id")

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "找不到用戶",
		})
		return
	}

	// 使用 DTO 返回數據
	userResponse := user.ToUserResponse()
	c.JSON(http.StatusOK, gin.H{
		"user": userResponse,
	})

}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// 獲取所有用戶
func GetUsers(c *gin.Context) {
	var users []models.User
	keyword := c.Query("keyword")

	db := database.DB
	if keyword != "" {
		like := "%" + keyword + "%"
		db = db.Where("username ILIKE ? OR email ILIKE ?", like, like)
	}

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "獲取用戶列表失敗",
		})
		return
	}

	// 轉換所有用戶數據為 DTO
	userResponses := make([]models.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = user.ToUserResponse()
	}

	c.JSON(http.StatusOK, gin.H{
		"users": userResponses,
	})
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
