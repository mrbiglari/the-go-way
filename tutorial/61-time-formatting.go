package main

import (
	"fmt"
	"time"
)

func timeFormatting() {
	print := fmt.Println
	time1 := time.Now()
	print(time1.Format(time.RFC3339))

	time2, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	print(time2)
	/*
		Go's time formatting uses the reference date:
		Mon Jan 2 15:04:05 MST 2006
		1   2  3  4  5    6
	*/
	print(time1.Format("3:04PM"))
	print(time1.Format("Mon Jan _2 15:04:05 2006"))
	format1 := "3 04 PM"
	time3, _ := time.Parse(format1, "8 50 AM")
	print(time3)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		time1.Year(), time1.Month(), time1.Day(),
		time1.Hour(), time1.Minute(), time1.Second())

	format2 := "Mon Jan _2 15:04:05 2006"
	_, _error := time.Parse(format2, "8:50 PM")
	print(_error)
}
