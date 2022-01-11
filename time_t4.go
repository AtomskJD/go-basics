package main

import (
	"fmt"
	"strings"
	"time"
)

const now = 1589570165

func main() {
	/*
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		strDur := sc.Text()
	*/
	strDur := "12 мин. 13 сек."
	strDurs := strings.Split(strDur, " ")
	dur, _ := time.ParseDuration(string(strDurs[0]) + `m` + string(strDurs[2]) + `s`)
	startDate := time.Unix(now, 0)
	startDate = startDate.UTC()

	fmt.Println(startDate.Add(dur).Format(time.UnixDate))
}
