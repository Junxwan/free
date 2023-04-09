package csv

import (
	"bytes"
	"fmt"
	file "github.com/Junxwan/free/lib"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type OPRawChipsCsv struct {
	df *dataframe.DataFrame
}

func NewOPRawChipsCsv(filePath string) (OPRawChipsCsv, error) {
	df, err := readCsv(filePath)
	if err != nil {
		return OPRawChipsCsv{}, fmt.Errorf("read path: %s csv error %w", filePath, err)
	}

	return OPRawChipsCsv{df: df}, nil
}

func (o OPRawChipsCsv) ToChipsCsv(outPath string) error {
	df := o.df.FilterAggregation(
		dataframe.And,
		dataframe.F{Colidx: 17, Comparator: series.Eq, Comparando: "一般"},
		dataframe.F{Colidx: 1, Comparator: series.Eq, Comparando: "TXO"},
	)

	records := df.Records()
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

		dir := filepath.Join(outPath, file.OpChipsDirPath(period))

		if err := saveCsv(dir, now, resultDf); err != nil {
			return fmt.Errorf("saveFile error: %w", err)
		}
	}

	return nil
}

func readCsv(filePath string) (*dataframe.DataFrame, error) {
	body, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile error: %w", err)
	}

	df := dataframe.ReadCSV(bytes.NewReader(body))
	return &df, nil
}
