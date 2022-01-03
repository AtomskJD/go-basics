package main

import (
	"fmt"
	"strings"
)

type Battery string

func (b Battery) String() string {
	s := string(b)
	tab := strings.Count(s, "0")
	pointers := strings.Count(s, "1")
	return "[" + strings.Repeat(" ", tab) + strings.Repeat("X", pointers) + "]"

}

func main() {
	var batterryForTest Battery
	fmt.Scan(&batterryForTest)
	fmt.Println(batterryForTest)
}
