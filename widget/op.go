package widget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Junxwan/free/csv"
	"github.com/Junxwan/free/data"
)

func NewDownLoadOPChipsButton(datePicker *DatePicker, parent fyne.Window) *widget.Button {
	return widget.NewButton("download OP Chips", func() {
		path := "E:\\我的雲端硬碟\\金融\\data"
		f, err := data.DownLoadOPChips(*datePicker.T, path)
		if err != nil {
			dialog.NewInformation("download", err.Error(), parent).Show()
			return
		}

		opCsv, err := csv.NewOPRawChipsCsv(f)
		if err != nil {
			dialog.NewInformation("download", fmt.Sprintf("load op chips raw csv error: %w", err.Error()), parent).Show()
			return
		}

		if err := opCsv.ToChipsCsv(path); err != nil {
			dialog.NewInformation("download", fmt.Sprintf("save op chips csv error: %w", err.Error()), parent).Show()
			return
		}

		dialog.NewInformation("Result", fmt.Sprintf("%s OP Chips Success", datePicker.T.Format("2006-01-02")), parent).Show()
	})
}
