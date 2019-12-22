package main

import "fmt"
import "time"

import "../../internal/apps/alm"

type Day struct {
        name string
	day int
        dow int // 0: sun, 1: mon ... 6: sat
        is_sat bool
        is_sun bool
        is_nhd bool
        memo string
}

type Month struct {
        name string
        day [31]Day
}

func main() {

	sdate := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 365; i++ {
		fmt.Printf("%s\n", sdate.String())
		sdate = sdate.Add(time.Duration(86400) * time.Second)
	}

	alm_view.Display_month()

	var day1 Day

        day1.day = 31
        day1.memo = "hoge"

	fmt.Print("%v", day1)

        fmt.Print("\n")

        var jan Month
        jan.day[1] = day1
        fmt.Print("%v", jan)
}
