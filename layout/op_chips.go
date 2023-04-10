package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	fwidget "github.com/Junxwan/free/widget"
	"image/color"
	"time"
)

func NewOPLayout(path string, now time.Time, win fyne.Window) *fyne.Container {
	dateText := canvas.NewText("DownLoad OP Chips:", color.White)
	dateText.TextSize = 15

	return container.New(layout.NewHBoxLayout(), dateText, fwidget.NewDownLoadOPChipsDatePicker(path, now, win).Button)
}
