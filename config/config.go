package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	ApiKey      string
	ApiSecret   string
	LogFile     string
	ProductCode string

	TradeDuration time.Duration
	Durations     map[string]time.Duration
	DbName        string
	SQLDriver     string
	Host          string
	Port          int

	BackTest         bool
	UsePercent       float64
	DataLimit        int
	StopLimitPercent float64
	NumRanking       int
	TradeBtcAmount   float64

	BotAccessToken       string
	ClientSecret         string
	SigningSecret        string
	TargetChannel        string
	UseSlackNotification bool

	ChannelSecret       string
	ChannelToken        string
	UserId              string
	UseLineNotification bool
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	durations := map[string]time.Duration{
		"1s": time.Second,
		"1m": time.Minute,
		"1h": time.Hour,
	}

	Config = ConfigList{
		ApiKey:           cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret:        cfg.Section("bitflyer").Key("api_secret").String(),
		LogFile:          cfg.Section("gotrading").Key("log_file").String(),
		ProductCode:      cfg.Section("gotrading").Key("product_code").String(),
		Durations:        durations,
		TradeDuration:    durations[cfg.Section("gotrading").Key("trade_duration").String()],
		DbName:           cfg.Section("db").Key("name").String(),
		SQLDriver:        cfg.Section("db").Key("driver").String(),
		Host:             cfg.Section("web").Key("host").String(),
		Port:             cfg.Section("web").Key("port").MustInt(),
		BackTest:         cfg.Section("gotrading").Key("back_test").MustBool(),
		UsePercent:       cfg.Section("gotrading").Key("use_percent").MustFloat64(),
		DataLimit:        cfg.Section("gotrading").Key("data_limit").MustInt(),
		StopLimitPercent: cfg.Section("gotrading").Key("stop_limit_percent").MustFloat64(),
		NumRanking:       cfg.Section("gotrading").Key("num_ranking").MustInt(),
		TradeBtcAmount:   cfg.Section("gotrading").Key("trade_btc_amount").MustFloat64(),

		BotAccessToken:       cfg.Section("slack").Key("bot_user_auth_access_token").String(),
		ClientSecret:         cfg.Section("slack").Key("client_secret").String(),
		SigningSecret:        cfg.Section("slack").Key("signing_secret").String(),
		TargetChannel:        cfg.Section("slack").Key("target_channel").String(),
		UseSlackNotification: cfg.Section("slack").Key("use_slack_notification").MustBool(),

		ChannelSecret:       cfg.Section("line").Key("channel_secret").String(),
		ChannelToken:        cfg.Section("line").Key("channel_token").String(),
		UserId:              cfg.Section("line").Key("user_id").String(),
		UseLineNotification: cfg.Section("line").Key("use_line_notification").MustBool(),
	}

	useEnv := os.Getenv("USE_ENV") == "true"
	if useEnv {
		Config.ApiKey = os.Getenv("ApiKey")
		Config.ApiSecret = os.Getenv("ApiSecret")
		Config.BotAccessToken = os.Getenv("BotAccessToken")
		Config.ClientSecret = os.Getenv("ClientSecret")
		Config.SigningSecret = os.Getenv("SigningSecret")
		Config.ChannelSecret = os.Getenv("ChannelSecret")
		Config.ChannelToken = os.Getenv("ChannelToken")
		Config.UserId = os.Getenv("UserId")

	}
}
