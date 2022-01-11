package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var sTime string
	// sTime = "2020-05-15 12:00:00"
	sTime, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	sTime = strings.Trim(sTime, "\n")
	// fmt.Scan(&sTime)
	t, err := time.Parse("2006-01-02 15:04:05", sTime)
	md := time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		12,
		00,
		00,
		0,
		t.Location(),
	)
	if err != nil {
		panic(err)
	}
	if t.After(md) {
		t = t.Add(time.Hour * 24)
	}
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}
