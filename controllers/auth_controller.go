package controllers

import (
	"fmt"
	"go-train/database"
	"go-train/models"
	"go-train/repositories"
	"go-train/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

// RegisterUser 處理用戶註冊
func RegisterUser(c *gin.Context) {

	var user models.User
	start := time.Now()

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, "無效的請求數據"))
		return
	}
	//驗證必要欄位
	if user.Username == "" || user.Password == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, "用戶名、密碼和郵箱都不能為空"))
		return
	}

	//加密密碼
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, "密碼加密失敗"))
		return
	}

	user.Password = string(hashedPassword)

	result := "success"
	if err := repositories.CreateUser(&user); err != nil {
		if strings.Contains(err.Error(), "unique constraint") || strings.Contains(err.Error(), "duplicate key") {
			result = "username_or_email_exists"
			c.JSON(http.StatusConflict, utils.ErrorResponse(http.StatusConflict, "用戶名或郵箱已被使用"))
		} else {
			result = "fail"
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, "註冊失敗"))
		}
		utils.WriteOperationLog(
			user.Username, "register", "", "", "", result, c.ClientIP(), "API", "User", "用戶註冊", time.Since(start).Milliseconds(),
		)
		return
	}

	// 清除返回數據中的密碼
	user.Password = ""
	userResponse := user.ToUserRegister()
	c.JSON(http.StatusOK, gin.H{"user": userResponse})

	utils.WriteOperationLog(
		user.Username, "register", fmt.Sprintf("%d", user.ID), "", "", "success", c.ClientIP(), "API", "User", "用戶註冊", time.Since(start).Milliseconds(),
	)

}

// LoginUser 處理用戶登入
func LoginUser(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	start := time.Now()

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, "無效的請求數據"))
		return
	}

	// 從數據庫獲取用戶
	user, err := repositories.GetUserByUsername(loginData.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse(http.StatusUnauthorized, "用戶名或密碼錯誤"))
		utils.WriteOperationLog(
			loginData.Username, "login", "", "", "", "fail", c.ClientIP(), "API", "User", "用戶登入失敗", time.Since(start).Milliseconds(),
		)
		return
	}

	// 驗證密碼
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse(http.StatusUnauthorized, "用戶名或密碼錯誤"))
		utils.WriteOperationLog(
			loginData.Username, "login", fmt.Sprintf("%d", user.ID), "", "", "fail", c.ClientIP(), "API", "User", "用戶登入失敗", time.Since(start).Milliseconds(),
		)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, "生成令牌失敗"))
		utils.WriteOperationLog(
			user.Username, "login", fmt.Sprintf("%d", user.ID), "", "", "fail", c.ClientIP(), "API", "User", "用戶登入失敗", time.Since(start).Milliseconds(),
		)
		return
	}

	// 清除返回數據中的密碼
	//user.Password = ""
	//c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "登入成功", user))

	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "登入成功", gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	}))

	utils.WriteOperationLog(
		user.Username, "login", fmt.Sprintf("%d", user.ID), "", "", "success", c.ClientIP(), "API", "User", "用戶登入", time.Since(start).Milliseconds(),
	)

}

func ChangePassword(c *gin.Context) {
	var req ChangePasswordRequest
	start := time.Now()
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "無效的數據請求",
		})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "未授權",
		})
		return
	}

	// 獲取用戶數據
	var user models.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "找不到用戶",
		})
		return
	}

	// 驗證舊密碼
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "舊密碼不正確"})
		utils.WriteOperationLog(
			user.Username, "change_password", fmt.Sprintf("%d", user.ID), "", "", "fail", c.ClientIP(), "API", "User", "用戶修改密碼失敗", time.Since(start).Milliseconds(),
		)
		return
	}

	// 加密新密碼
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密碼加密失敗"})
		utils.WriteOperationLog(
			user.Username, "change_password", fmt.Sprintf("%d", user.ID), "", "", "fail", c.ClientIP(), "API", "User", "用戶修改密碼失敗", time.Since(start).Milliseconds(),
		)
		return
	}

	// 更新密碼
	user.Password = string(hashedPassword)
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新密碼失敗"})
		utils.WriteOperationLog(
			user.Username, "change_password", fmt.Sprintf("%d", user.ID), "", "", "fail", c.ClientIP(), "API", "User", "用戶修改密碼失敗", time.Since(start).Milliseconds(),
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "密碼已成功更新",
	})

	utils.WriteOperationLog(
		user.Username, "change_password", fmt.Sprintf("%d", user.ID), "", "", "success", c.ClientIP(), "API", "User", "用戶修改了密碼", time.Since(start).Milliseconds(),
	)

}
