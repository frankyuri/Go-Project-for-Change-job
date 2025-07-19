package utils

import (
	"os"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	os.Setenv("JWT_SECRET_KEY", "testsecret")
	token, err := GenerateToken(1, "testuser")
	if err != nil {
		t.Fatalf("產生 token 失敗: %v", err)
	}
	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("驗證 token 失敗: %v", err)
	}
	if claims.UserID != 1 || claims.Username != "testuser" {
		t.Errorf("claims 不正確: %+v", claims)
	}
}
