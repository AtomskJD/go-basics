package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now.Year())
	date := time.Date(
		2020,
		time.May,
		10,
		12,
		00,
		00,
		1,
		time.UTC,
	)
	fmt.Println(date)
	fmt.Println(now.Unix())

	fmt.Println(now.Format("02/01/2006 Mon 15:04:05"))
	fmt.Println(date.Format("02 Jan 2006 Mon"))
	loc, err := time.LoadLocation("Asia/Yekaterinburg")
	if err != nil {
		panic(err)
	}
	fmt.Println(loc)
}
