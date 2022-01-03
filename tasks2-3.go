package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	for i := 9; i > 0; i-- {
		if strings.Contains(s, strconv.Itoa(i)) {
			fmt.Println(i)
			break
		}
	}

}
