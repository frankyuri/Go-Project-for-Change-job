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
