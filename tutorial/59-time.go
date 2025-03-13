package main

import (
	"fmt"
	"time"
)

func _time() {
	print := fmt.Println
	now := time.Now()
	print(now) // 2023-10-05 15:04:05.123456789 +0000 UTC m=+0.000000001

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	print(then) // 2009-11-17 20:34:58.651387237 +0000 UTC

	print(then.Year(), then.Month(), then.Day(),
		then.Hour(), then.Minute(), then.Second(), then.Nanosecond()) // 2009 November 17 20 34 58 651387237

	print(then.Location()) // UTC
	print(then.Weekday())  // Tuesday

	print(then.Before(now), then.After(now), then.Equal(now)) // true false false

	difference := now.Sub(then)
	print(difference) // 123456h7m8.123456789s
	print(int(difference.Hours()/(24*365)), int(difference.Hours()/24), int(difference.Hours()/(24*30)),
		difference.Hours(), difference.Minutes(), difference.Seconds(), difference.Nanoseconds()) // 15 5594 186 134270.80075711556 8.056248045426934e+06 4.8337488272561604e+08 483374882725616063

	print(then.Add(difference))  // 2023-10-05 15:04:05.123456789 +0000 UTC
	print(then.Add(-difference)) // 1995-01-30 01:05:52.179317685 +0000 UTC
}
