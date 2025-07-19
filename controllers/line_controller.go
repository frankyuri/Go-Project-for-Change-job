package controllers

import (
	"go-train/database"
	"go-train/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func LineWebhook(c *gin.Context) {
    var req struct {
        Events []struct {
            Type   string `json:"type"`
            Source struct {
                UserID string `json:"userId"`
            } `json:"source"`
            Message struct {
                Type string `json:"type"`
                Text string `json:"text"`
            } `json:"message"`
        } `json:"events"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    for _, event := range req.Events {
        if event.Type == "message" && event.Message.Type == "text" {
            content := strings.TrimSpace(event.Message.Text)
            status := "todo"
            if strings.HasPrefix(content, "/done") {
                status = "done"
                content = strings.TrimSpace(strings.TrimPrefix(content, "/done"))
            }
            todo := models.LineTodo{
                UserID:    event.Source.UserID,
                Content:   content,
                Status:    status,
                CreatedAt: time.Now(),
                UpdatedAt: time.Now(),
            }
            database.DB.Create(&todo)
        }
    }

    c.JSON(http.StatusOK, gin.H{"status": "ok"})
}