package data

import (
	"bytes"
	"fmt"
	file "github.com/Junxwan/free/lib"
	"github.com/go-gota/gota/dataframe"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"os"
	"path/filepath"
	"time"
)

// 下載期交所OP未平倉資料
func DownLoadOPChips(time time.Time, outPath string) (string, error) {
	f := filepath.Join(outPath, file.OpRawChipsPathByDay(time))

	if !file.IsExist(f) {
		body, err := downLoadOPChips(time)
		if err != nil {
			return "", fmt.Errorf("downLoadOPChips error: %w", err)
		}

		if err := saveOPRawChips(body, f); err != nil {
			return "", fmt.Errorf("saveOPRawChips error: %w", err)
		}
	}

	return f, nil
}

func saveOPRawChips(body []byte, path string) error {
	// 因為期交所csv每個row結尾都多一個,會跟column數量對不上出現wrong number of fields
	// 解決方式就是在多一個空column
	body = bytes.Replace(body, []byte("漲跌%"), []byte("漲跌%,"), 1)
	df := dataframe.ReadCSV(bytes.NewReader(body))

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
