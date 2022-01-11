package main

import (
	"fmt"
	"time"
)

func main() {
	var sTime string
	// sTime := "1986-04-16T05:20:00+06:00"
	fmt.Scan(&sTime)
	t, err := time.Parse("2006-01-02T15:04:05-07:00", sTime)
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Format(time.UnixDate))
}
