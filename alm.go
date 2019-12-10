package main

import "fmt"
import "time"

func main() {
    sdate := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
    for i := 0; i < 365; i++ {
        fmt.Printf("%s\n", sdate.String())
        sdate = sdate.Add(time.Duration(86400) * time.Second)
    }
}
