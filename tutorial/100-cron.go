package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

/*
cron expression:
* * * * *
│ │ │ │ │
│ │ │ │ └── Day of week (0–6, Sun–Sat)
│ │ │ └──── Month (1–12)
│ │ └────── Day of month (1–31)
│ └──────── Hour (0–23)
└────────── Minute (0–59)
*/

func _cron() {
	c := cron.New()

	c.AddFunc("@every 5s", func() { fmt.Println("1") })       // every 5 seconds
	c.AddFunc("0 9 * * *", func() { fmt.Println("2") })       // at 9:00 AM every day
	c.AddFunc("*/5 * * * *", func() { fmt.Println("3") })     // every 5 minutes
	c.AddFunc("3/5 * * * *", func() { fmt.Println("4") })     // every 5 minutes starting at minute 3 (i.e., 3, 8, 13, ...)
	c.AddFunc("0/5 15 * * *", func() { fmt.Println("5") })    // every 5 minutes during hour 15 (i.e., 15:00, 15:05, ..., 15:55)
	c.AddFunc("0/5 15-23 * * *", func() { fmt.Println("6") }) // every 5 minutes from 15:00 to 23:55

	c.Start()

	time.Sleep(10 * time.Second)

	c.Stop()
}
