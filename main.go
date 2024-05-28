package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type K struct {
	Time time.Time
	S1   string
	S2   string
	D    int64
	O    int64
	H    int64
	L    int64
	C    int64
	V    int64
}

var data map[string][]K

func main() {
	data = make(map[string][]K)

	load("5")
	load("30")
	load("60")
	loadDay()
	loadWeek()
	loadMonth()

	router := gin.New()

	router.Static("/public", "./public")

	var Time struct {
		T     string `form:"t"`
		Start int64  `form:"start"`
		End   int64  `form:"end"`
		Is    int64  `form:"is"`
	}

	router.GET("/kline", func(c *gin.Context) {
		c.ShouldBindQuery(&Time)

		if Time.T == "" {
			Time.T = "5"
		}

		vv := data[Time.T]

		var min, max int

		d := [][]int64{}
		if Time.End > 0 {
			switch Time.T {
			case "day":
				d = getDay(Time.End, Time.Is)
			case "week":
				d = getWeek(Time.End, Time.Is)
			case "month":
				d = getMonth(Time.End, Time.Is)
			}

			c.JSON(http.StatusOK, d)

			return
		} else {
			min = 0
			max = len(vv)
			switch Time.T {
			case "5":
				min = max - 1200
			case "30":
				min = max - 1200
			case "60":
				min = max - 1200
			}
		}

		for _, v := range vv[min:max] {
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

func load(t string) {
	d := make([]K, 0, 10000)
	y := time.Now().Year()
	for i := 0; i < 10; i++ {
		file, err := os.Open(fmt.Sprintf("./data/tfe-tx00-%d-%smin.csv", y-i, t))
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
			vv, _ := strconv.ParseInt(v[7], 10, 64)

			s := strings.Split(v[2], " ")

			d = append(d, K{
				Time: t,
				S1:   s[0],
				S2:   s[1],
				D:    t.UnixMilli(),
				O:    o,
				H:    h,
				L:    l,
				C:    c,
				V:    vv,
			})
		}
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i].D < d[j].D
	})

	data[t] = d
}

func loadDay() {
	var start int
	var open K

	vv := data["60"]

	data["day_o"] = []K{}

	for i, v := range vv {
		// 開始有夜盤
		if v.S1 >= "2017-5-15" {
			if v.S2 == "16:00:00" {
				open = v
				start = i
			}
		} else {
			if v.S2 == "09:45:00" {
				open = v
				start = i
			}
		}

		if v.S2 == "05:00:00" {
			t, _ := time.Parse("2006-01-02", vv[i+1].S1)

			k := K{
				Time: t,
				S1:   vv[i+1].S1,
				S2:   v.S2,
				D:    t.UnixMilli(),
				O:    open.O,
				C:    v.C,
				H:    v.H,
				L:    v.L,
			}

			d := vv[start : i+1]
			for _, v := range d {
				k.V = k.V + v.V

				if v.H > k.H {
					k.H = v.H
				}

				if v.L < k.L {
					k.L = v.L
				}
			}

			data["day_o"] = append(data["day_o"], k)
		}

		if v.S2 == "13:45:00" {
			t, _ := time.Parse("2006-01-02", v.S1)

			k := K{
				Time: t,
				S1:   v.S1,
				S2:   v.S1,
				D:    t.UnixMilli(),
				O:    open.O,
				C:    v.C,
				H:    v.H,
				L:    v.L,
			}

			d := vv[start : i+1]
			for _, v := range d {
				k.V = k.V + v.V

				if v.H > k.H {
					k.H = v.H
				}

				if v.L < k.L {
					k.L = v.L
				}
			}

			data["day"] = append(data["day"], k)
		}
	}
}

func loadWeek() {
	var start int
	var open K

	vv := data["day"]

	for i, v := range vv {
		if v.Time.Weekday() == time.Monday {
			open = v
			start = i
		}

		if v.Time.Weekday() == time.Friday {
			t, _ := time.Parse("2006-01-02", v.S1)

			k := K{
				Time: t,
				S1:   open.S1,
				S2:   v.S1,
				D:    t.UnixMilli(),
				O:    open.O,
				C:    v.C,
				H:    v.H,
				L:    v.L,
			}

			d := vv[start : i+1]
			for _, v := range d {
				//k.V = k.V + v.V

				if v.H > k.H {
					k.H = v.H
				}

				if v.L < k.L {
					k.L = v.L
				}
			}

			data["week"] = append(data["week"], k)
		}
	}
}

func loadMonth() {
	var start int
	last := data["day"][0].S1[:7]

	vv := data["day"]
	for i, v := range vv {
		if v.S1[:7] != last {
			last = v.S1[:7]

			t, _ := time.Parse("2006-01-02", v.S1)

			k := K{
				Time: t,
				S1:   vv[start].S1,
				S2:   vv[i-1].S1,
				D:    t.UnixMilli(),
				O:    vv[start].O,
				C:    vv[i-1].C,
				H:    v.H,
				L:    v.L,
			}

			d := vv[start:i]
			for _, v := range d {
				//k.V = k.V + v.V

				if v.H > k.H {
					k.H = v.H
				}

				if v.L < k.L {
					k.L = v.L
				}
			}

			data["month"] = append(data["month"], k)

			start = i
		}
	}
}

func getDay(end int64, is int64) [][]int64 {
	vv := data["day"]
	d := [][]int64{}

	for _, v := range vv {
		if v.D <= end {
			d = append(d, []int64{
				v.D, v.O, v.H, v.L, v.C, v.V,
			})
		}
	}

	if is == 1 {
		v := d[len(d)-1]
		t := time.UnixMilli(v[0])
		s := t.Format("2006-01-02")

		for _, vv := range data["day_o"] {
			if vv.S1 == s {
				d[len(d)-1] = []int64{
					vv.D, vv.O, vv.H, vv.L, vv.C, vv.V,
				}
			}
		}

		return d
	}

	return d
}

func getWeek(end int64, is int64) [][]int64 {
	return [][]int64{}
}

func getMonth(end int64, is int64) [][]int64 {
	return [][]int64{}
}
