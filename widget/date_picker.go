package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	xWidget "fyne.io/x/fyne/widget"
	"time"
)

func NewDatePicker(now time.Time, win fyne.Window) *widget.Button {
	datePicker := widget.NewButton(now.Format("2006-01-02"), nil)

	calendar := xWidget.NewCalendar(now, func(t2 time.Time) {
		datePicker.Text = t2.Format("2006-01-02")
		datePicker.Refresh()
	})

	datePicker.OnTapped = func() {
		dialog.ShowCustom("Date", "Cancel", calendar, win)
	}

	return datePicker
}
