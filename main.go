package main

import (
	"flag"
	"fmt"
	"time"
)

const (
	secondsaday = 3600 * 24
)

func main() {
	dateFlag := flag.Int("date", 0, "A date, given as unix days")
	flag.Parse()
	if *dateFlag != 0 {
		fmt.Println(time.Unix(int64(*dateFlag)*secondsaday, 0).UTC().Format("02.01.2006"))
		return
	}
	fmt.Println(time.Now().Unix() / secondsaday)
}
