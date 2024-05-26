package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"
)

type K struct {
	Time time.Time
	D    int64
	O    int64
	H    int64
	L    int64
	C    int64
	V    int64
}

func main() {
	var data []K

	y := time.Now().Year()
	for i := 0; i < 2; i++ {
		file, err := os.Open(fmt.Sprintf("./data/tfe-tx00-%d-5min.csv", y-i))
		if err != nil {
			panic(err)
		}

		defer file.Close()

		df := csv.NewReader(file)
		csvs, err := df.ReadAll()
		if err != nil {
			panic(err)
		}

		for _, v := range csvs {
			t, err := time.Parse("2006-01-02 15:04:05", v[2])
			if err != nil {
				panic(err)
			}

			o, _ := strconv.ParseInt(v[3], 10, 64)
			h, _ := strconv.ParseInt(v[4], 10, 64)
			l, _ := strconv.ParseInt(v[5], 10, 64)
			c, _ := strconv.ParseInt(v[6], 10, 64)
			v, _ := strconv.ParseInt(v[7], 10, 64)

			data = append(data, K{
				Time: t,
				D:    t.UnixMilli(),
				O:    o,
				H:    h,
				L:    l,
				C:    c,
				V:    v,
			})
		}
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].D < data[j].D
	})

	router := gin.New()

	router.Static("/public", "./public")

	var Time struct {
		Start int64 `form:"start"`
		End   int64 `form:"end"`
	}

	router.GET("/kline", func(c *gin.Context) {
		c.ShouldBindQuery(&Time)

		var min, max int

		d := [][]int64{}
		if Time.Start > 0 {
			for i, v := range data {
				if min == 0 && Time.Start <= v.D {
					min = i
				}

				if min != 0 && Time.End <= v.D {
					max = i
					break
				}
			}
		} else {
			min = len(data) - 360
			max = len(data)
		}

		for _, v := range data[min:max] {
			d = append(d, []int64{
				v.D, v.O, v.H, v.L, v.C, v.V,
			})
		}

		c.JSON(http.StatusOK, d)
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.Run()
}
