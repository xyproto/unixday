package main

import (
	"flag"
	"fmt"
	"time"
)

const (
	// UNIX time does not take leap seconds into account. Neither does Unix Days.
	secondsaday = 3600 * 24
)

func main() {
	dateFlag := flag.Int("date", 0, "A date, given as unix days")
	flag.Parse()
	if *dateFlag != 0 {
		// Be aware that UNIX days and normal days does not match, because of leap seconds.
		fmt.Println(time.Unix(int64(*dateFlag)*secondsaday, 0).UTC().Format("02.01.2006"))
		return
	}
	fmt.Println(time.Now().Unix() / secondsaday)
}
