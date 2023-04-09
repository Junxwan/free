package widget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Junxwan/free/data"
)

func NewDownLoadOPChipsButton(datePicker *DatePicker, parent fyne.Window) *widget.Button {
	return widget.NewButton("download OP Chips", func() {
		err := data.DownLoadOPChips(*datePicker.T, "E:\\我的雲端硬碟\\金融\\data")
		message := fmt.Sprintf("%s OP Chips Success", datePicker.T.Format("2006-01-02"))
		if err != nil {
			message = err.Error()
		}

		dialog.NewInformation("Result", message, parent).Show()
	})
}
