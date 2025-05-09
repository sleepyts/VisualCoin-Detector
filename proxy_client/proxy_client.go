package proxy_client

import (
	"net/http"
	"net/url"
)

var Client *http.Client

func InitClient() {
	proxyUrl, _ := url.Parse("http://127.0.0.1:7897")
	Client = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
}
