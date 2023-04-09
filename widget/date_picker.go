package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	xWidget "fyne.io/x/fyne/widget"
	"time"
)

type DatePicker struct {
	T      *time.Time
	Button *widget.Button
}

func NewDatePicker(now time.Time, win fyne.Window) *DatePicker {
	datePicker := &DatePicker{
		T:      &now,
		Button: widget.NewButton(now.Format("2006-01-02"), nil),
	}

	calendar := xWidget.NewCalendar(now, func(t2 time.Time) {
		datePicker.T = &t2
		datePicker.Button.Text = t2.Format("2006-01-02")
		datePicker.Button.Refresh()
	})

	datePicker.Button.OnTapped = func() {
		dialog.ShowCustom("Date", "Cancel", calendar, win)
	}

	return datePicker
}
