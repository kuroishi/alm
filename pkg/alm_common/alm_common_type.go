package alm_common_type

import "fmt"
import "time"

type Day struct {
     	Date   time.Time
	name   string
	Day    int
	dow    int // 0: sun, 1: mon ... 6: sat
	is_sat bool
	is_sun bool
	is_nhd bool
	memo   string
}

type Month struct {
	name string
	Day  []Day
}

func debug() {
     fmt.Print("x")
}