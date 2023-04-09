package file

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func OpRawChipsPathByDay(time time.Time) string {
	return fmt.Sprintf("%s/%s.csv",
		filepath.Join("op", "day"),
		time.Format("2006-01-02"),
	)
}

func OpChipsDirPath(name string) string {
	return filepath.Join("op", "chips", name)
}

func OpChipsPath(time time.Time, dirName string) string {
	return fmt.Sprintf("%s/%s.csv",
		OpChipsDirPath(dirName),
		time.Format("2006-01-02"),
	)
}
