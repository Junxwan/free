package layout

import (
	"github.com/Junxwan/free/csv"
	"testing"
)

func Test_makeOPChipsTableData(t *testing.T) {
	csv.Init("E:\\我的雲端硬碟\\金融\\data")

	f, _ := csv.ReadOPChipsByPeriod("202304W2")
	makeOPChipsTableData(f)
}
