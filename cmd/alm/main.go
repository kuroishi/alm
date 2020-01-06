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

	var mon []Month
	mon = make([]Month, 3)
	mon[0].Day = make([]Day, 31)
	mon[1].Day = make([]Day, 31)
	mon[2].Day = make([]Day, 31)
	fmt.Print("\nmon:\n")
	//now := time.Now()
	init_month(mon[0], time.Date(1976, 9, 1, 0, 0, 0, 0, time.Local))
	init_month(mon[1], time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))
	alm_view.Display_month(mon)
	fmt.Print("\n")
}

/*

	
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
    "github.com/PuerkitoBio/goquery"
)

func ScrapeScrape() bool {
    return true
}

func Scrape(w string) {
    query := fmt.Sprintf("https://eow.alc.co.jp/search?q=%s", w)
    res, err := http.Get(query)
    if err != nil {
        log.Fatal(err)
    }
    doc, err := goquery.NewDocumentFromReader(res.Body);

    //html, err := doc.Find("body").Html()
    //fmt.Print(html)

    num, _ := strconv.Atoi(doc.Find("#itemsNumber > strong").Text())
    if num == 0 {
        os.Exit(1)
    }

    word := doc.Find("#resultsList > ul:nth-child(3) > li:nth-child(1) > span:nth-child(1) > h2:nth-child(1) > span:nth-child(1)").Text()
    word_class := doc.Find("#resultsList > ul:nth-child(3) > li:nth-child(1) .wordclass").Text()
    fmt.Print(word)
    fmt.Print(" | ")
    fmt.Print(word_class)
    fmt.Print("\n----\n")
    ScrapeScrape()
    doc.Find("#resultsList > ul:nth-child(3) > li:nth-child(1) > div:nth-child(2) > ol:nth-child(2) > li").Each(func(idx int, s *goquery.Selection) {
             idx++
             fmt.Printf("%d\n", idx)
             fmt.Printf("%s\n\n", s.Text())
         })

    doc.Find("#resultsList > ul:nth-child(3) > li:nth-child(1) > div:nth-child(2)").Each(func(idx int, s *goquery.Selection) {
        s.Find("li").Each(func(idx int, s *goquery.Selection) {
            fmt.Printf("%+v\n", s)
        })
    })

}

func main() {
    if len(os.Args) != 2 {
        os.Exit(1)
    }
    //fmt.Printf("args: %s", os.Args[1])
    Scrape(os.Args[1])
}
oke-office-vm00-dev%

*/
