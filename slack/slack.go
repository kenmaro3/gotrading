package slack

import (
	"fmt"

	"gotrading/config"

	"gotrading/app/models"

	"github.com/slack-go/slack"
)

func Message(message string, s *models.SignalEvent) {
	tradeContent := fmt.Sprintf("Trade details: %v", s)
	api := slack.New(config.Config.BotAccessToken)
	attachment := slack.Attachment{
		Pretext: tradeContent,
		Text:    "",
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}

	channelID, timestamp, err := api.PostMessage(
		"test",
		slack.MsgOptionText(message, false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Println("Message successfully sent to channel %s at %s", channelID, timestamp)
}
