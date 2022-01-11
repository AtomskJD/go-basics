package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	/*
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		str := sc.Text()
	*/
	str := "13.03.2018 14:00:15,12.03.2018 14:00:15"
	strs := strings.Split(str, ",")
	startTime, _ := time.Parse("02.01.2006 15:04:05", strs[0])
	endTime, _ := time.Parse("02.01.2006 15:04:05", strs[1])
	var dur time.Duration
	if startTime.After(endTime) {
		dur = startTime.Sub(endTime)
	} else {
		dur = endTime.Sub(startTime)
	}
	fmt.Println(dur)
}
