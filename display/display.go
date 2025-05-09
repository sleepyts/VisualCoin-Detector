package display

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var App fyne.App
var MainWindow fyne.Window
var PriceLabel *widget.Label

var PriceLabels []*widget.Label

func Init() {

	// 创建 Fyne 应用
	App = app.New()
	MainWindow = App.NewWindow("Tracker")

	PriceLabel = widget.NewLabel("Loading...\n")

	// 设置窗口内容
	MainWindow.SetContent(container.NewVBox(
		PriceLabel,
	))
}
