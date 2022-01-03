package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	_, err := fmt.Scan(&a, &b)
	if err == nil {
		res := a*a + b*b
		fmt.Println(math.Sqrt(res))
	}
}
