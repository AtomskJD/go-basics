package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	ss := strings.Split(s, "")
	fmt.Println(strings.Join(ss, "*"))
}
