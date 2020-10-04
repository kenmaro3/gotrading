package line

import (
	"fmt"

	"gotrading/app/models"
	"gotrading/config"

	"github.com/line/line-bot-sdk-go/linebot"
)

func Message(message string, s *models.SignalEvent) {
	tradeContent := fmt.Sprintf("%s \n Trade details: %v", message, s)
	bot, err := linebot.New(config.Config.ChannelSecret, config.Config.ChannelToken)
	if err != nil {
		fmt.Println(err)
	}
	if _, err := bot.PushMessage(config.Config.UserId, linebot.NewTextMessage(tradeContent)).Do(); err != nil {
		fmt.Println(err)
	}
}
