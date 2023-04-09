package csv

import (
	"fmt"
	file "github.com/Junxwan/free/lib"
	"github.com/go-gota/gota/dataframe"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"os"
	"path/filepath"
	"time"
)

func saveFile(path string, df dataframe.DataFrame) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Create path: %s error: %w", path, err)
	}

	defer f.Close()

	writer := transform.NewWriter(f, unicode.UTF8BOM.NewEncoder())
	if err := df.WriteCSV(writer, dataframe.WriteHeader(true)); err != nil {
		return fmt.Errorf("WriteCSV error: %w", err)
	}

	return nil
}

func saveCsv(dir string, time time.Time, df dataframe.DataFrame) error {
	if !file.IsExist(dir) {
		if err := os.Mkdir(dir, 0750); err != nil {
			return fmt.Errorf("mkdir path: %s error: %w", dir, err)
		}
	}

	if err := saveFile(filepath.Join(dir, time.Format("2006-01-02")+".csv"), df); err != nil {
		return fmt.Errorf("saveFile error: %w", err)
	}
	return nil
}
