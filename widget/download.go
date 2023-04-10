package widget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/Junxwan/free/csv"
	"github.com/Junxwan/free/data"
	"time"
)

// 下載OP未平倉
func NewDownLoadOPChipsDatePicker(path string, now time.Time, win fyne.Window) *DatePicker {
	return NewDatePickerOnSelected(now, win, func(selected time.Time) {
		f, err := data.DownLoadOPChips(selected, path)
		if err != nil {
			dialog.NewInformation("download", err.Error(), win).Show()
			return
		}

		if err := csv.ToChipsCsv(f, path); err != nil {
			dialog.NewInformation("download", fmt.Sprintf("save op chips csv error: %w", err.Error()), win).Show()
			return
		}

		dialog.NewInformation("Result", fmt.Sprintf("%s OP Chips Success", selected.Format("2006-01-02")), win).Show()
	})
}
