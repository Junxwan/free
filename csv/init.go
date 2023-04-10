package csv

import (
	"fmt"
	"path/filepath"
	"time"
)

var OPChipsDirPath string
var OPRawChipsPathByDay string

func init() {
	OPChipsDirPath = filepath.Join("op", "chips")
	OPRawChipsPathByDay = filepath.Join("op", "day")
}

func OpChipsDirPath(name string) string {
	return filepath.Join(OPChipsDirPath, name)
}

func OpRawChipsPathByDay(time time.Time) string {
	return fmt.Sprintf("%s/%s.csv", OPRawChipsPathByDay, time.Format("2006-01-02"))
}
