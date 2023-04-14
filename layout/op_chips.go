package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Junxwan/free/csv"
	fwidget "github.com/Junxwan/free/widget"
	"github.com/go-gota/gota/dataframe"
	"image/color"
	"sort"
	"strconv"
	"time"
)

func NewOPLayout(now time.Time, win fyne.Window) *fyne.Container {
	dateText := canvas.NewText("DownLoad OP Chips:", color.White)
	dateText.TextSize = 15

	btn := widget.NewButton("View", func() {
		opw := fyne.CurrentApp().NewWindow("OP Chips View")
		opw.SetContent(newOPListLayout(win))
		opw.Resize(fyne.NewSize(1500, 900))
		opw.Show()
	})

	return container.New(layout.NewHBoxLayout(), dateText, fwidget.NewDownLoadOPChipsDatePicker(now, win).Button, btn)
}

//	C    	周別    	  P
//
// 日期(近->遠)  履約價  日期(近->遠)
// 未平倉(增減)
func newOPListLayout(win fyne.Window) *fyne.Container {
	periods, _ := csv.Periods()
	opChipsPeriod, err := csv.ReadOPChipsByPeriod(periods[0])
	if err != nil {
		dialog.NewInformation("load op chips", err.Error(), win).Show()
		return nil
	}

	tableData := makeOPChipsTableData(opChipsPeriod)
	table := newOPChipsTable(tableData)
	column := newOPChipsTableColumn(tableData)

	selectPeriods := widget.NewSelect(periods, func(period string) {
		opChipsPeriod, err := csv.ReadOPChipsByPeriod(period)
		if err != nil {
			dialog.NewInformation("load op chips", err.Error(), win).Show()
		}

		*tableData = *makeOPChipsTableData(opChipsPeriod)

		table.Refresh()
		column.Refresh()
	})

	selectPeriods.Selected = periods[0]

	periodLabel := canvas.NewText("period: ", color.White)
	periodLabel.TextSize = 15

	return container.NewBorder(container.NewVBox(container.NewHBox(periodLabel, selectPeriods), column), nil, nil, nil, table)
}

// op 未平倉table column
func newOPChipsTableColumn(data *opChipsTable) *widget.Table {
	table := widget.NewTable(
		func() (int, int) { return 1, len(data.columns) },
		func() fyne.CanvasObject {
			test := canvas.NewText("", color.White)
			test.Alignment = fyne.TextAlignCenter
			test.TextStyle = fyne.TextStyle{Bold: true}
			return test
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*canvas.Text)
			switch id.Row {
			case 0:
				label.TextSize = 17
				label.Text = data.columns[id.Col]
			}
		})

	for i := 0; i < len(data.columns)+1; i++ {
		table.SetColumnWidth(i, 100)
	}

	return table
}

// op 未平倉table data
func newOPChipsTable(data *opChipsTable) *widget.Table {
	table := widget.NewTable(
		func() (int, int) { return len(data.prices), (len(data.columns) * 2) - 1 },
		func() fyne.CanvasObject {
			test := canvas.NewText("", color.White)
			test.Alignment = fyne.TextAlignCenter
			test.TextStyle = fyne.TextStyle{Bold: true}
			return test
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*canvas.Text)
			price := data.prices[id.Row]
			test := ""

			if id.Col == (data.center)*2 {
				label.TextSize = 17
				test = price
			} else {
				label.TextSize = 15
				var v int
				var date string

				if id.Col > 12 {
					date = data.dates[(id.Col-1)/2]
				} else {
					date = data.dates[id.Col/2]
				}

				switch id.Col {
				// C
				case 1, 3, 5, 7, 9, 11:
					v = data.GetCall(date, price)
					if v > 10000 {
						label.Color = &color.RGBA{R: 0xFF, G: 0, B: 0, A: 150}
					}

				// C diff
				case 0, 2, 4, 6, 8, 10:
					v = data.GetCallDiff(date, price)
					if v < 0 {
						label.Color = &color.RGBA{R: 0, G: 0xC0, B: 0x3E, A: 100}
					} else {
						label.Color = &color.RGBA{R: 0xFF, G: 0, B: 0, A: 150}
					}

				// P
				case 13, 15, 17, 19, 21, 23:
					v = data.GetPut(date, price)
					if v > 10000 {
						label.Color = &color.RGBA{R: 0xFF, G: 0, B: 0, A: 150}
					}

				// P diff
				case 14, 16, 18, 20, 22, 24:
					v = data.GetPutDiff(date, price)
					if v < 0 {
						label.Color = &color.RGBA{R: 0, G: 0xC0, B: 0x3E, A: 100}
					} else {
						label.Color = &color.RGBA{R: 0xFF, G: 0, B: 0, A: 150}
					}
				}

				test = strconv.Itoa(v)
			}

			label.Text = test
		})

	for i := 0; i < (len(data.columns)*2)-1; i++ {
		if i == len(data.columns)-1 {
			table.SetColumnWidth(i, 100)
		} else {
			table.SetColumnWidth(i, 47)
		}

	}

	for i := 0; i < len(data.prices); i++ {
		table.SetRowHeight(i, 20)
	}

	return table
}

type opChipsTable struct {
	dates   []string
	columns []string
	prices  []string
	data    map[string]map[string][]string

	center int
}

func (o *opChipsTable) GetCall(date, price string) int {
	return o.Get(1, date, price)
}

func (o *opChipsTable) GetPut(date, price string) int {
	return o.Get(2, date, price)
}

func (o *opChipsTable) GetCallDiff(date, price string) int {
	return o.GetDiff(1, date, price)
}

func (o *opChipsTable) GetPutDiff(date, price string) int {
	return o.GetDiff(2, date, price)
}

func (o *opChipsTable) Get(index int, date, price string) int {
	value, ok := o.data[date][price]
	if !ok {
		return 0
	}

	v, _ := strconv.Atoi(value[index])

	return v
}

func (o *opChipsTable) GetDiff(index int, date, price string) int {
	v := o.Get(index, date, price)
	bv := v

	for i, d := range o.dates {
		if date == d && i > 0 {
			bv = o.Get(index, o.dates[i-1], price)
			break
		}
	}

	return v - bv
}

func makeOPChipsTableData(opRawChips map[string]*dataframe.DataFrame) *opChipsTable {
	dates := []string{}
	data := make(map[string]map[string][]string)
	pricesMap := make(map[string]int)

	for date, value := range opRawChips {
		data[date] = make(map[string][]string)

		dates = append(dates, date)
		records := value.Records()
		for _, v := range records[1:len(records)] {
			data[date][v[0]] = v

			if _, ok := pricesMap[v[0]]; !ok {
				pricesMap[v[0]] = 0
			}
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(dates)))
	if len(dates) > 6 {
		dates = dates[:6]
	} else {
		for i := 0; i <= 7-len(dates); i++ {
			dates = append(dates, "")
		}
	}

	sort.Sort(sort.StringSlice(dates))
	dst := make([]string, len(dates))
	copy(dst, dates)
	sort.Sort(sort.Reverse(sort.StringSlice(dst)))
	dates = append(dates, dst...)

	prices := []string{}
	for k, _ := range pricesMap {
		prices = append(prices, k)
	}

	sort.Sort(sort.StringSlice(prices))

	if len(prices) > 30 {
		pl := int(len(prices) / 2)
		prices = append(prices[pl-15:pl], prices[pl:pl+15]...)
	}

	opChipsTable := &opChipsTable{
		dates:  dates,
		prices: prices,
		data:   data,
	}

	dst = make([]string, len(dates))
	copy(dst, dates)
	opChipsTable.center = len(dates) / 2
	opChipsTable.columns = append(dst[:len(dst)/2], append([]string{"Price"}, dst[len(dst)/2:]...)...)

	return opChipsTable
}
