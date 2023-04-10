package csv

import (
	"fmt"
	"path/filepath"
	"time"
)

const (
	OP_DIR_NAME           = "op"
	OP_CHIPS_DIR_NAME     = "chips"
	OP_CHIPS_DAY_DIR_NAME = "day"
)

var path string
var opDirPath string
var opChipsDirPath string
var opRawChipsPathByDay string

func Init(dataPath string) {
	path = dataPath
	opDirPath = filepath.Join(path, OP_DIR_NAME)
	opChipsDirPath = filepath.Join(opDirPath, OP_CHIPS_DIR_NAME)
	opRawChipsPathByDay = filepath.Join(opDirPath, OP_CHIPS_DAY_DIR_NAME)

}

func GetOpChipsPath() string {
	return opChipsDirPath
}

func GetOpChipsPathByPeriod(name string) string {
	return filepath.Join(opChipsDirPath, name)
}

func GetOpRawChipsPathByDay(time time.Time) string {
	return fmt.Sprintf("%s/%s.csv", opRawChipsPathByDay, time.Format("2006-01-02"))
}
