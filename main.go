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
	t()
	//t1()
}

func t1() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text3 := canvas.NewText("(right)", color.White)
	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	text4 := canvas.NewText("centered", color.White)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	myWindow.ShowAndRun()
}

func t() {
	a := app.New()
	win := a.NewWindow("期權")

	dateText := canvas.NewText("Date:", color.White)
	dateText.TextSize = 20

	content := container.New(layout.NewHBoxLayout(), dateText, fwidget.NewDatePicker(time.Now(), win))

	win.SetContent(container.New(layout.NewVBoxLayout(), content))

	win.Resize(fyne.NewSize(720, 480))
	win.ShowAndRun()
}
