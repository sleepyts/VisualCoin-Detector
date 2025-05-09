package main

import (
	"encoding/json"
	"io"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"gihub.com/sleepyts/Visual_Coin_Detector/display"
	"gihub.com/sleepyts/Visual_Coin_Detector/model"
	"gihub.com/sleepyts/Visual_Coin_Detector/proxy_client"
)

func main() {
	proxy_client.InitClient()

	display.Init()

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				baseUrl := "https://www.okx.com"
				path := "/api/v5/market/ticker?instId=ETH-USDT"
				resp, err := proxy_client.Client.Get(baseUrl + path)
				if err != nil {
					log.Println("Error making request:", err)
					continue
				}

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Println("Error reading response body:", err)
					continue
				}

				var response model.ApiResponse
				err = json.Unmarshal(body, &response)
				if err != nil {
					log.Println("Error unmarshalling JSON:", err)
					continue
				}

				fyne.Do(func() {
					// 更新标签内容
					display.PriceLabel.SetText("ETH-USDT: " + response.Data[0].Last)
					display.PriceLabel.Refresh() // 刷新标签以显示新数据
				})
			}
		}
	}()

	// 启动应用
	display.MainWindow.ShowAndRun()
}
