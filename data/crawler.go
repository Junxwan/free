package data

import (
	"fmt"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"net/url"
	"time"
)

func downLoadOPChips(time time.Time) ([]byte, error) {
	date := time.Format("2006/01/02")

	resp, err := http.PostForm("https://www.taifex.com.tw/cht/3/optDataDown", url.Values{
		"down_type":      []string{"1"},
		"commodity_id":   []string{"TXO"},
		"commodity_id2":  []string{""},
		"queryStartDate": []string{date},
		"queryEndDate":   []string{date},
	})

	if err != nil {
		return nil, fmt.Errorf("http.PostForm error: %w", err)
	}

	defer resp.Body.Close()

	chips, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)

	}

	result, _, err := transform.Bytes(traditionalchinese.Big5.NewDecoder(), chips)
	if err != nil {
		return nil, fmt.Errorf("transform.Bytes error: %w", err)
	}

	return result, nil
}
