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
	for i := 0; i < 15; i++ {
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

		if v.S2 >= "13:30:00" {
			t, _ := time.Parse("2006-01-02", v.S1)

			if t.Weekday() == time.Wednesday {
				if v.S2 != "13:30:00" {
					continue
				}
			} else if v.S2 != "13:45:00" {
				continue
			}

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
	size := len(vv)

	open = vv[0]

	for i, v := range vv {
		if size > i+1 && !isSameWeek(time.UnixMilli(v.D), time.UnixMilli(vv[i+1].D)) {
			t, _ := time.Parse("2006-01-02", v.S1)

			k := K{
				Time: t,
				S1:   open.S1,
				S2:   v.S1,
				D:    open.D,
				O:    open.O,
				C:    v.C,
				H:    v.H,
				L:    v.L,
			}

			d := vv[start : i+1]
			for _, v := range d {
				if v.H > k.H {
					k.H = v.H
				}

				if v.L < k.L {
					k.L = v.L
				}
			}

			data["week"] = append(data["week"], k)

			open = vv[i+1]
			start = i + 1
		}
	}

	k := K{
		Time: open.Time,
		S1:   open.S1,
		S2:   vv[size-1].S2,
		D:    open.D,
		O:    open.O,
		C:    vv[size-1].C,
		H:    open.H,
		L:    open.L,
	}

	d := vv[start:]
	for _, v := range d {
		if v.H > k.H {
			k.H = v.H
		}

		if v.L < k.L {
			k.L = v.L
		}
	}

	data["week"] = append(data["week"], k)
}

func loadMonth() {
	var start int
	var open K

	vv := data["day"]
	size := len(vv)

	open = vv[0]

	for i, v := range vv {
		if size > i+1 && !isSameMonth(time.UnixMilli(v.D), time.UnixMilli(vv[i+1].D)) {
			t, _ := time.Parse("2006-01-02", v.S1)

			k := K{
				Time: t,
				S1:   open.S1,
				S2:   v.S1,
				D:    open.D,
				O:    open.O,
				C:    v.C,
				H:    v.H,
				L:    v.L,
			}

			d := vv[start : i+1]
			for _, v := range d {
				if v.H > k.H {
					k.H = v.H
				}

				if v.L < k.L {
					k.L = v.L
				}
			}

			data["month"] = append(data["month"], k)

			open = vv[i+1]
			start = i + 1
		}
	}

	k := K{
		Time: open.Time,
		S1:   open.S1,
		S2:   vv[size-1].S2,
		D:    open.D,
		O:    open.O,
		C:    vv[size-1].C,
		H:    open.H,
		L:    open.L,
	}

	d := vv[start:]
	for _, v := range d {
		if v.H > k.H {
			k.H = v.H
		}

		if v.L < k.L {
			k.L = v.L
		}
	}

	data["month"] = append(data["month"], k)
}

func getDay(end int64, is int64) [][]int64 {
	vv := data["day"]
	d := [][]int64{}
	t := time.UnixMilli(end).Format("2006-01-02")

	for _, v := range vv {
		if v.S1 <= t {
			d = append(d, []int64{
				v.D, v.O, v.H, v.L, v.C, v.V,
			})
		} else {
			break
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
	vv := data["week"]
	kk := []K{}
	t := time.UnixMilli(end).Format("2006-01-02")

	for _, v := range vv {
		if v.S1 <= t {
			kk = append(kk, v)
		}
	}

	endi := len(kk) - 1

	dd := []K{}
	for _, v := range data["day"] {
		if vv[endi].S1 <= v.S1 && v.S1 <= t {
			dd = append(dd, v)
		}
	}

	l := len(kk)

	if is == 1 {
		for _, v := range data["day_o"] {
			if v.S1 == t {
				kk[l-1].C = v.C
			}
		}
	} else {
		kk[l-1].C = dd[len(dd)-1].C
	}

	kk[l-1].S2 = t

	for _, v := range dd {
		if v.H > kk[l-1].H {
			kk[l-1].H = v.H
		}

		if v.L < kk[l-1].L {
			kk[l-1].L = v.L
		}
	}

	d := [][]int64{}
	for _, k := range kk {
		d = append(d, []int64{
			k.D, k.O, k.H, k.L, k.C, k.V,
		})
	}

	return d
}

func getMonth(end int64, is int64) [][]int64 {
	vv := data["month"]
	kk := []K{}
	t := time.UnixMilli(end).Format("2006-01-02")

	for _, v := range vv {
		if v.S1 <= t {
			kk = append(kk, v)
		}
	}

	endi := len(kk) - 1

	dd := []K{}
	for _, v := range data["day"] {
		if vv[endi].S1 <= v.S1 && v.S1 <= t {
			dd = append(dd, v)
		}
	}

	l := len(kk)

	if is == 1 {
		for _, v := range data["day_o"] {
			if v.S1 == t {
				kk[l-1].C = v.C
			}
		}
	} else {
		kk[l-1].C = dd[len(dd)-1].C
	}

	kk[l-1].S2 = t

	for _, v := range dd {
		if v.H > kk[l-1].H {
			kk[l-1].H = v.H
		}

		if v.L < kk[l-1].L {
			kk[l-1].L = v.L
		}
	}

	d := [][]int64{}
	for _, k := range kk {
		d = append(d, []int64{
			k.D, k.O, k.H, k.L, k.C, k.V,
		})
	}

	return d
}

func isSameWeek(date1, date2 time.Time) bool {
	startOfWeek1 := getStartOfWeek(date1)
	startOfWeek2 := getStartOfWeek(date2)
	return startOfWeek1.Equal(startOfWeek2)
}

func isSameMonth(date1, date2 time.Time) bool {
	return date1.Year() == date2.Year() && date1.Month() == date2.Month()
}

func getStartOfWeek(date time.Time) time.Time {
	weekday := int(date.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	startOfWeek := date.AddDate(0, 0, -weekday+1)
	return startOfWeek
}
