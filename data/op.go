package data

import (
	"bytes"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"os"
	"path/filepath"
	"time"
)

func DownLoadOPChips(time time.Time, path string) error {
	chips, err := downLoadOPChips(time)
	if err != nil {
		return err
	}

	// 因為期交所csv每個row結尾都多一個,會跟column數量對不上出現wrong number of fields
	// 解決方式就是在多一個空column
	chips = bytes.Replace(chips, []byte("漲跌%"), []byte("漲跌%,"), 1)
	df := dataframe.ReadCSV(bytes.NewReader(chips))

	path = fmt.Sprintf("%s/%s.csv",
		filepath.Join(path, "op", "day"),
		time.Format("2006-01-02"),
	)

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
