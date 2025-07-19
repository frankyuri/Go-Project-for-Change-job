package controllers

import (
	"fmt"
	"go-train/database"
	"go-train/models"
	"go-train/utils"
	"net/http"
	"strconv"
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
			switch {
			case strings.HasPrefix(content, "/todo"):
				// 新增
				todo := models.LineTodo{
					UserID:  event.Source.UserID,
					Content: strings.TrimSpace(strings.TrimPrefix(content, "/todo")),
					Status:  "todo",
				}
				database.DB.Create(&todo)
				utils.ReplyToUser(event.Source.UserID, "已新增待辦事項: "+todo.Content)
			case strings.HasPrefix(content, "/done"):
				// 完成
				id := strings.TrimSpace(strings.TrimPrefix(content, "/done"))
				if id == "" {
					utils.ReplyToUser(event.Source.UserID, "請輸入 /done <id>")
					break
				}
				result := database.DB.Model(&models.LineTodo{}).
					Where("id = ? AND user_id = ?", id, event.Source.UserID).
					Update("status", "done")
				if result.RowsAffected == 0 {
					utils.ReplyToUser(event.Source.UserID, "找不到該任務或已完成")
				} else {
					utils.ReplyToUser(event.Source.UserID, "已標記完成")
				}
			case strings.HasPrefix(content, "/edit"):
				// 編輯
				args := strings.Fields(strings.TrimPrefix(content, "/edit"))
				if len(args) < 2 {
					utils.ReplyToUser(event.Source.UserID, "格式錯誤，請輸入 /edit <id> <新內容>")
					break
				}
				id := args[0]
				newContent := strings.Join(args[1:], " ")
				if _, err := strconv.Atoi(id); err != nil {
					utils.ReplyToUser(event.Source.UserID, "請輸入正確的 id")
					break
				}
				result := database.DB.Model(&models.LineTodo{}).
					Where("id = ? AND user_id = ?", id, event.Source.UserID).
					Update("content", newContent)
				if result.RowsAffected == 0 {
					utils.ReplyToUser(event.Source.UserID, "找不到該任務")
				} else {
					utils.ReplyToUser(event.Source.UserID, "已更新內容")
				}
			case strings.HasPrefix(content, "/show"):
				now := time.Now()
				startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
				endOfDay := startOfDay.Add(24 * time.Hour)

				var todos []models.LineTodo
				database.DB.Where("user_id = ? AND created_at >= ? AND created_at < ?", event.Source.UserID, startOfDay, endOfDay).Find(&todos)

				var todoList, doneList []string
				for _, t := range todos {
					line := fmt.Sprintf("%d. %s", t.ID, t.Content)
					if t.Status == "todo" {
						todoList = append(todoList, line)
					} else {
						doneList = append(doneList, line)
					}
				}
				reply := "【待辦事項】\n"
				if len(todoList) == 0 {
					reply += "無\n"
				} else {
					reply += strings.Join(todoList, "\n") + "\n"
				}
				reply += "【已完成】\n"
				if len(doneList) == 0 {
					reply += "無"
				} else {
					reply += strings.Join(doneList, "\n")
				}
				utils.ReplyToUser(event.Source.UserID, reply)
			default:
				utils.ReplyToUser(event.Source.UserID, "未知指令，請輸入 /todo, /done, /edit, /show")

			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
