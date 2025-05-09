package config

import (
	"log"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

type Config struct {
	// Whether to use proxy or not
	UseProxy bool `ini:"use_proxy"`
	// Proxy configuration
	Proxy ProxyConfig
	// Coins to track
	Coins []CoinConfig

	BaseApiUrl string `ini:"base_api_url"`
}

type ProxyConfig struct {
	// URL of the proxy server
	ProxyURL string `ini:"proxy_url"`
	// Port of the proxy server
	ProxyPort int `ini:"proxy_port"`
}

var AppConfig *Config

type CoinConfig struct {
	CoinName string `ini:"coin_name"`
}

func InitConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		if len(os.Args) > 1 {
			cfg, _ = ini.Load(os.Args[1])
		} else {
			log.Panicf("Failed to load config file: %v", err)
		}
	}
	AppConfig = &Config{}
	AppConfig.UseProxy = cfg.Section("proxy").Key("use_proxy").MustBool(false)

	cfg.Section("proxy").MapTo(&AppConfig.Proxy)
	coinStr := cfg.Section("coins").Key("coin_name").String()
	coinNames := strings.Split(coinStr, ",")
	for _, c := range coinNames {
		AppConfig.Coins = append(AppConfig.Coins, CoinConfig{CoinName: strings.TrimSpace(c)})
	}

	AppConfig.BaseApiUrl = cfg.Section("api").Key("base_api_url").String()
}

func GetBaseApiUrl() string {
	return AppConfig.BaseApiUrl
}

func UseProxy() bool {
	return AppConfig.UseProxy
}
