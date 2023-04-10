package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/Junxwan/free/csv"
	flayout "github.com/Junxwan/free/layout"
	"time"
)

func main() {
	a := app.New()
	win := a.NewWindow("期權")
	csv.Init("E:\\我的雲端硬碟\\金融\\data")

	now := time.Now()

	win.SetContent(container.New(layout.NewVBoxLayout(), flayout.NewOPLayout(now, win)))

	win.Resize(fyne.NewSize(720, 480))
	win.ShowAndRun()
}
