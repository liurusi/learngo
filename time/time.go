package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now() // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%04d\n", t.Day(), t.Month(), t.Year())
	//date.month.year
	t = time.Now().UTC() // Wed Dec 21 08:52:14 +0000 UTC 2011
	fmt.Println(t)
	fmt.Println(time.Now())
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 //must be in nanosec
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
	// formatting times:
	fmt.Println(t.Format(time.RFC822))         //21 Dec 11 0852 UTC
	fmt.Println(t.Format(time.ANSIC))          // Wed Dec 21 08:56:34 2011
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
	s := t.Format("20060102")
	fmt.Println(t, "=>", s) // Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221

}
