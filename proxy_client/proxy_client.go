package proxy_client

import (
	"net/http"
	"net/url"

	"gihub.com/sleepyts/Visual_Coin_Detector/config"
)

var Client *http.Client

func InitClient() {
	if config.UseProxy() {
		proxyUrl, _ := url.Parse("http://127.0.0.1:7897")
		Client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
	} else {
		Client = &http.Client{}
	}

}
