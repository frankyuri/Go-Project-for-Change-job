package dto

import (
	"go-train/models"
)

type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserRegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// 添加一個轉換方法
func ToUserResponse(u *models.User) UserResponse {
	return UserResponse{
		ID:        u.DisplayID,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToUserRegisterResponse(u *models.User) UserRegisterResponse {
	return UserRegisterResponse{
		ID:       u.DisplayID,
		Username: u.Username,
		Email:    u.Email,
	}
}
