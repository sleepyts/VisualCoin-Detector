package proxy_client

import (
	"net/http"
	"net/url"
	"strconv"

	"gihub.com/sleepyts/Visual_Coin_Detector/config"
)

var Client *http.Client

func InitClient() {
	if config.UseProxy() {
		proxyUrl, _ := url.Parse(config.AppConfig.Proxy.ProxyURL + ":" + strconv.Itoa(config.AppConfig.Proxy.ProxyPort))
		Client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
	} else {
		Client = &http.Client{}
	}

}
