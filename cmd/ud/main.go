package main

import (
	"flag"
	"fmt"

	"github.com/xyproto/unixday"
)

func main() {
	dateFlag := flag.Int("date", -1, "A date, given as UNIX days")
	flag.Parse()
	if *dateFlag != -1 {
		fmt.Println(unixday.UNIXDay(*dateFlag).Date())
		return
	}
	fmt.Println(unixday.Now())
}
