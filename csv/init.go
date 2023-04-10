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

var _path string
var _opDirPath string
var _opChipsDirPath string
var _opRawChipsPathByDay string

func Init(dataPath string) {
	_path = dataPath
	_opDirPath = filepath.Join(_path, OP_DIR_NAME)
	_opChipsDirPath = filepath.Join(_opDirPath, OP_CHIPS_DIR_NAME)
	_opRawChipsPathByDay = filepath.Join(_opDirPath, OP_CHIPS_DAY_DIR_NAME)

}

func GetOpChipsPath() string {
	return _opChipsDirPath
}

func GetOpChipsPathByPeriod(name string) string {
	return filepath.Join(_opChipsDirPath, name)
}

func GetOpRawChipsPathByDay(time time.Time) string {
	return fmt.Sprintf("%s/%s.csv", _opRawChipsPathByDay, time.Format("2006-01-02"))
}
