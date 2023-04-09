package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	fwidget "github.com/Junxwan/free/widget"
	"image/color"
	"time"
)

func main() {
	a := app.New()
	win := a.NewWindow("期權")
	path := "E:\\我的雲端硬碟\\金融\\data"
	now := time.Now()

	dateText := canvas.NewText("DownLoad OP Chips:", color.White)
	dateText.TextSize = 20

	dateContent := container.New(layout.NewHBoxLayout(), dateText, fwidget.NewDownLoadOPChipsDatePicker(path, now, win).Button)

	win.SetContent(container.New(layout.NewVBoxLayout(), dateContent))

	win.Resize(fyne.NewSize(720, 480))
	win.ShowAndRun()
}
