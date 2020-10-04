package main

import (
	"fmt"

	"gotrading/config"
	"gotrading/slack"
)

func main() {
	// utils.LoggingSettings(config.Config.LogFile)
	// controllers.StreamIngestionData()
	// log.Println(controllers.StartWebServer())

	fmt.Println(config.Config.BotAccessToken)
	slack.Message()
}
