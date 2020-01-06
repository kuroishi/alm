package main

import "fmt"
import "time"
//import "reflect"

import "../../internal/apps/alm"
import . "../../pkg/alm_common"

func init_month(m Month, tm time.Time) int {
	t := time.Date(tm.Year(), tm.Month(), 1, 0, 0, 0, 0, time.Local)
	fmt.Print(t.Month().String() + "\n")
	fmt.Print("xxx\n")
	nextm := t.AddDate(0, 1, 0)
	sentinel := time.Date(nextm.Year(), nextm.Month(), 1, 0, 0, 0, 0, time.Local)
	for sentinel.Equal(t) == false {
	        fmt.Print("---")
		fmt.Print(t)
		fmt.Print("\n")
		m.Day[t.Day()-1].Date = t
		t = t.AddDate(0, 0, 1)
	}

	return 1
}

func main() {
	sdate := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 365; i++ {
		fmt.Printf("%s\n", sdate.String())
		sdate = sdate.Add(time.Duration(86400) * time.Second)
	}

	t := time.Now()
	fmt.Print(t)

	var mon Month
	mon.Day = make([]Day, 31)
	fmt.Print("\nmon:\n")
	//now := time.Now()
	init_month(mon, time.Date(1976, 9, 1, 0, 0, 0, 0, time.Local))
	alm_view.Display_month(mon)
	fmt.Print("\n")
}
