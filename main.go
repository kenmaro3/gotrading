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

	// fmt.Println(config.Config.ApiKey)
	// fmt.Println(config.Config.ChannelToken)

}

// package main

// import (
// 	"time"

// 	"github.com/kmihara/gotrading/app/models"
// )

// func main() {
// 	s := models.NewSignalEvents()
// 	df, _ := models.GetAllCandle("BTC_JPY", time.Minute, 10)
// 	c1 := df.Candles[0]
// 	c2 := df.Candles[1]
// 	s.Buy("BTC_JPY", c1.Time.UTC(), c1.Close, 1.0, true)
// 	s.Sell("BTC_JPY", c2.Time.UTC(), c2.Close, 1.0, true)

// }

// func main(){
// 	api := bitflyer.New(
// 		config.Config.ApiKey,
// 		config.Config.ApiSecret,
// 	)
// 	order := &bitflyer.Order{
// 		ProductCode:     "BTC_JPY",
// 		ChildOrderType:  "MARKET",
// 		Side:            "SELL",
// 		Size:            0.001,
// 		MinuteToExpires: 10000,
// 		TimeInForce:     "GTC",
// 	}

// 	_, err := api.SendOrder(order)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
