package main

import (
	"encoding/json"
	"io"
	"log"
	"os/exec"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"gihub.com/sleepyts/Visual_Coin_Detector/config"
	"gihub.com/sleepyts/Visual_Coin_Detector/display"
	"gihub.com/sleepyts/Visual_Coin_Detector/model"
	"gihub.com/sleepyts/Visual_Coin_Detector/proxy_client"
)

func SetAlwaysOnTop(windowTitle string) {
	exec.Command("wmctrl", "-r", windowTitle, "-b", "add,above").Run()
}
func main() {

	config.InitConfig()

	proxy_client.InitClient()

	display.Init()

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				baseUrl := config.GetBaseApiUrl()
				showLabel := ""
				for index := range config.AppConfig.Coins {
					coinName := config.AppConfig.Coins[index].CoinName
					path := "/api/v5/market/ticker?instId=" + coinName + "-USDT"
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

					last, _ := strconv.ParseFloat(response.Data[0].Last, 64)
					sodUtc0, _ := strconv.ParseFloat(response.Data[0].SodUtc0, 64)
					rate := (last - sodUtc0) / sodUtc0 * 100

					buyPrice := config.AppConfig.Coins[index].BuyPrice
					buyNum := config.AppConfig.Coins[index].BuyNum
					profit := (last - buyPrice) * buyNum

					showLabel += coinName + "-USDT: " + response.Data[0].Last + "   " + strconv.FormatFloat(rate, 'f', 2, 64) + "%\n" + "Profit: " + strconv.FormatFloat(profit, 'f', 2, 64) + " USDT\n"

				}
				fyne.Do(func() {
					display.PriceLabel.SetText(showLabel)
					display.PriceLabel.Refresh()
				})
			}
		}
	}()

	// 启动应用
	display.MainWindow.Show()

	SetAlwaysOnTop(display.MainWindow.Title())

	display.App.Run()

}
