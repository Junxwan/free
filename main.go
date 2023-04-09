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

	dateText := canvas.NewText("Date:", color.White)
	dateText.TextSize = 20

	datePicker := fwidget.NewDatePicker(time.Now(), win)
	dateContent := container.New(layout.NewHBoxLayout(), dateText, datePicker.Button)
	downloadOPBtn := fwidget.NewDownLoadOPChipsButton(datePicker, win)

	tidyContent := container.New(layout.NewHBoxLayout(), downloadOPBtn)

	win.SetContent(container.New(layout.NewVBoxLayout(), dateContent, tidyContent))

	win.Resize(fyne.NewSize(720, 480))
	win.ShowAndRun()
}
