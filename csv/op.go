package csv

import (
	"bytes"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func ToChipsCsv(filePath, outPath string) error {
	df, err := readCsv(filePath)
	if err != nil {
		return fmt.Errorf("read path: %s csv error %w", filePath, err)
	}

	records := df.FilterAggregation(
		dataframe.And,
		dataframe.F{Colidx: 17, Comparator: series.Eq, Comparando: "一般"},
		dataframe.F{Colidx: 1, Comparator: series.Eq, Comparando: "TXO"},
	).Records()

	periodMap := make(map[string]map[string][][]string)

	now, err := time.Parse("2006/01/02", records[1][0])
	if err != nil {
		return fmt.Errorf("time.Parse date: %s error: %w", records[1][0], err)
	}

	for _, value := range records[1:len(records)] {
		period := strings.Trim(value[2], " ")
		if _, ok := periodMap[period]; !ok {
			periodMap[period] = make(map[string][][]string)
		}

		price := strings.Split(value[3], ".")[0]

		periodMap[period][price] = append(periodMap[period][price], value)
	}

	for period, rows := range periodMap {
		keys := make([]string, 0, len(rows))
		for k := range rows {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		csv := [][]string{[]string{"Price", "C", "P"}}
		for _, k := range keys {
			var call, put string

			if rows[k][0][4] == "買權" {
				call = rows[k][0][11]
			}
			if rows[k][1][4] == "買權" {
				call = rows[k][1][11]
			}
			if rows[k][0][4] == "賣權" {
				put = rows[k][0][11]
			}
			if rows[k][1][4] == "賣權" {
				put = rows[k][1][11]
			}

			csv = append(csv, []string{k, call, put})
		}

		resultDf := dataframe.LoadRecords(csv, dataframe.WithTypes(map[string]series.Type{
			"Price": series.String,
			"C":     series.Int,
			"P":     series.Int,
		}))

		if err := saveCsv(GetOpChipsPathByPeriod(period), now, resultDf); err != nil {
			return fmt.Errorf("saveFile error: %w", err)
		}
	}

	return nil
}

func ReadOPChipsByPeriod(period, path string) (map[string]*dataframe.DataFrame, error) {
	dirPath := filepath.Join(path, period)
	fs, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("os.ReadDir error: %w", err)
	}

	sort.Slice(fs, func(i, j int) bool {
		return fs[i].Name() < fs[j].Name()
	})

	dfs := make(map[string]*dataframe.DataFrame)
	for _, f := range fs {
		name := strings.Split(f.Name(), ".")[0]
		csv, err := readCsv(filepath.Join(dirPath, f.Name()))
		if err != nil {
			return nil, fmt.Errorf("readCsv _path: %s error: %w", filepath.Join(dirPath, f.Name()), err)
		}

		dfs[name] = csv
	}

	return dfs, nil
}

func Periods() ([]string, error) {
	fs, err := os.ReadDir(_opChipsDirPath)
	if err != nil {
		return nil, fmt.Errorf("os.ReadDir error: %w", err)
	}

	var periods []string
	for _, f := range fs {
		periods = append(periods, f.Name())
	}

	sort.Sort(sort.Reverse(sort.StringSlice(periods)))

	return periods, nil
}

func readCsv(filePath string) (*dataframe.DataFrame, error) {
	body, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile error: %w", err)
	}

	df := dataframe.ReadCSV(bytes.NewReader(body))
	return &df, nil
}
