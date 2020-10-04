package main

import (
	"log"

	"github.com/kmihara/gotrading/app/controllers"
	"github.com/kmihara/gotrading/config"
	"github.com/kmihara/gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())

}
