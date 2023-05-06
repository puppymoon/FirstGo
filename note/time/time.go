package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	//2023-05-03 09:18:02.489113 +0800 CST m=+0.000117143
	fmt.Println(t)
	//03.05.2023
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	// 2023-05-03 09:18:02.48911 +0800
	fmt.Println(t.Format("2006-01-02 15:04:05.99999 -0700"))
	// 03 May 23 09:18.48911 +0800
	fmt.Println(t.Format("02 Jan 06 15:04.99999 -0700"))
	// 03 May 23 09:18 CST
	fmt.Println(t.Format(time.RFC822))
	// Wed May  3 09:18:02 2023
	fmt.Println(t.Format(time.ANSIC))
	// 9:18AM
	fmt.Println(t.Format(time.Kitchen))

	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())

	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now)

	s := t.Format("20060102")
	// 2023-05-03 01:18:02.489509 +0000 UTC => 20230503
	fmt.Println(t, "=>", s)
}
