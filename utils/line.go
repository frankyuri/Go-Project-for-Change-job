package utils

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var Bot *linebot.Client

func InitLineBot(channelSecret, channelToken string) error {
	var err error
	Bot, err = linebot.New(channelSecret, channelToken)
	return err
}

func ReplyToUser(userID, message string) {
	if Bot != nil {
		_, err := Bot.PushMessage(userID, linebot.NewTextMessage(message)).Do()
		if err != nil {
			fmt.Println("LINE推送失敗:", err)
		}
	}
}

func ReplyHelpFlex(userID string) {
	flexJson := `{
  "type": "bubble",
  "header": {
    "type": "box",
    "layout": "vertical",
    "contents": [
      {
        "type": "text",
        "text": "📖 指令教學",
        "weight": "bold",
        "size": "lg"
      }
    ]
  },
  "body": {
    "type": "box",
    "layout": "vertical",
    "spacing": "md",
    "contents": [
      {
        "type": "box",
        "layout": "baseline",
        "contents": [
          {
            "type": "text",
            "text": "📝 新增待辦",
            "flex": 2,
            "size": "xs"
          },
          {
            "type": "text",
            "text": "/todo 內容",
            "flex": 3,
            "color": "#888888",
            "size": "xs"
          }
        ]
      },
      {
        "type": "box",
        "layout": "baseline",
        "contents": [
          {
            "type": "text",
            "text": "✅ 完成待辦",
            "flex": 2,
            "size": "xs"
          },
          {
            "type": "text",
            "text": "/done id",
            "flex": 3,
            "color": "#888888",
            "size": "xs"
          }
        ]
      },
      {
        "type": "box",
        "layout": "baseline",
        "contents": [
          {
            "type": "text",
            "text": "✏️ 編輯待辦",
            "flex": 2,
            "size": "xs"
          },
          {
            "type": "text",
            "text": "/edit id 新內容",
            "flex": 3,
            "color": "#888888",
            "size": "xs"
          }
        ]
      },
      {
        "type": "box",
        "layout": "baseline",
        "contents": [
          {
            "type": "text",
            "text": "📋 顯示今日待辦",
            "flex": 2,
            "size": "xs"
          },
          {
            "type": "text",
            "text": "/show",
            "flex": 3,
            "color": "#888888",
            "size": "xs"
          }
        ]
      },
      {
        "type": "box",
        "layout": "baseline",
        "contents": [
          {
            "type": "text",
            "text": "ℹ️ 顯示說明",
            "flex": 2,
            "size": "xs"
          },
          {
            "type": "text",
            "text": "/help",
            "flex": 3,
            "color": "#888888",
            "size": "xs"
          }
        ]
      }
    ]
  }
}`
	container, err := linebot.UnmarshalFlexMessageJSON([]byte(flexJson))
	if err != nil {
		fmt.Println("FlexMessage 解析失敗:", err)
		return
	}
	if Bot != nil {
		_, err := Bot.PushMessage(userID, linebot.NewFlexMessage("指令教學", container)).Do()
		if err != nil {
			fmt.Println("LINE推送失敗:", err)
		}
	}
}
